package aws

import (
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	endpointsBatchSize = 50
)

func (c *Component) loadUpdaters() {
	go func() {
		ticker := time.NewTicker(c.config.GetDuration(ConfigAwsUpdaterApplicationsDuration))

		for {
			select {
			case <-ticker.C:
				c.updaterApplications()
			case <-c.applicationsRun:
				c.updaterApplications()
			case d := <-c.applicationsTicker:
				ticker = time.NewTicker(d)
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(c.config.GetDuration(ConfigAwsUpdaterSubscriptionsDuration))

		for {
			select {
			case <-ticker.C:
				c.updaterSubscriptions()
			case <-c.subscriptionsRun:
				c.updaterSubscriptions()
			case d := <-c.subscriptionsTicker:
				ticker = time.NewTicker(d)
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(c.config.GetDuration(ConfigAwsUpdaterTopicsDuration))

		for {
			select {
			case <-ticker.C:
				c.updaterTopics()
			case <-c.topicsRun:
				c.updaterTopics()
			case d := <-c.topicsTicker:
				ticker = time.NewTicker(d)
			}
		}
	}()

	if c.config.GetBool(ConfigAwsRunUpdatersOnStartup) {
		c.applicationsRun <- true
		c.subscriptionsRun <- true
		c.topicsRun <- true
	}
}

func (c *Component) updaterApplications() {
	lastUpdate := time.Now().UTC()
	applications := map[string]AwsSnsApplication{}
	params := &sns.ListPlatformApplicationsInput{}

	err := c.GetSNS().ListPlatformApplicationsPages(params, func(p *sns.ListPlatformApplicationsOutput, lastPage bool) bool {
		for _, a := range p.PlatformApplications {
			arn := sdk.StringValue(a.PlatformApplicationArn)

			app := AwsSnsApplication{
				Arn:                   arn,
				AwsAttributes:         a.Attributes,
				Enabled:               true,
				EndpointsCount:        -1,
				EndpointsEnabledCount: -1,
				LastUpdate:            lastUpdate,
			}

			if dateRaw, ok := a.Attributes["AppleCertificateExpirationDate"]; ok {
				if dateValue, err := time.Parse(time.RFC3339, sdk.StringValue(dateRaw)); err == nil {
					app.CertificateExpirationDate = &dateValue
				}
			}

			if dateRaw, ok := a.Attributes["Enabled"]; ok && sdk.StringValue(dateRaw) == "false" {
				app.Enabled = false
			}

			applications[arn] = app
		}

		return !lastPage
	})

	if err != nil {
		c.logger.Errorf("Update applications error %s", err.Error())
		return
	}

	if metricApplicationsTotal != nil {
		metricApplicationsTotal.Set(float64(len(applications)))
	}

	c.mutex.Lock()
	for arn, application := range applications {
		if exists, ok := c.applications[arn]; ok {
			application.EndpointsCount = exists.EndpointsCount
			application.EndpointsEnabledCount = exists.EndpointsEnabledCount

			applications[arn] = application
		}
	}

	c.applications = applications
	c.mutex.Unlock()

	c.logger.Debugf("Updater found %d applications", len(applications))

	c.updaterEndpoints()
}

func (c *Component) updaterEndpoints() {
	applications := c.GetApplications()
	batchStartIndex := 0

	for i := range applications {
		params := &sns.ListEndpointsByPlatformApplicationInput{
			PlatformApplicationArn: sdk.String(applications[i].Arn),
		}

		applications[i].EndpointsCount = 0
		applications[i].EndpointsEnabledCount = 0
		applications[i].LastUpdate = time.Now().UTC()

		err := c.GetSNS().ListEndpointsByPlatformApplicationPages(params, func(p *sns.ListEndpointsByPlatformApplicationOutput, lastPage bool) bool {
			applications[i].EndpointsCount += len(p.Endpoints)

			for _, point := range p.Endpoints {
				if enabled, ok := point.Attributes["Enabled"]; ok && sdk.StringValue(enabled) == "true" {
					applications[i].EndpointsEnabledCount++
				}
			}

			return !lastPage
		})

		if err == nil {
			if metricEndpointsTotal != nil {
				metricEndpointsTotal.With("arn", applications[i].Arn).Set(float64(applications[i].EndpointsCount))
			}

			if metricEndpointsEnabled != nil {
				metricEndpointsEnabled.With("arn", applications[i].Arn).Set(float64(applications[i].EndpointsEnabledCount))
			}
		}

		if err != nil {
			c.logger.Errorf("Update apn %s is failed", applications[i].Arn, map[string]interface{}{
				"application.ednpoints":         applications[i].EndpointsCount,
				"application.ednpoints-enabled": applications[i].EndpointsEnabledCount,
				"error": err.Error(),
			})
		} else {
			c.logger.Debugf("Update apn %s is success", applications[i].Arn, map[string]interface{}{
				"application.ednpoints":         applications[i].EndpointsCount,
				"application.ednpoints-enabled": applications[i].EndpointsEnabledCount,
			})
		}

		// flush data
		batchEndIndex := i + 1

		if batchEndIndex%endpointsBatchSize == 0 || batchEndIndex == len(applications) {
			c.mutex.Lock()
			for _, current := range applications[batchStartIndex:batchEndIndex] {
				last, ok := c.applications[current.Arn]
				if !ok {
					continue
				}

				last.EndpointsCount = current.EndpointsCount
				last.EndpointsEnabledCount = current.EndpointsEnabledCount
				last.LastUpdate = current.LastUpdate

				c.applications[current.Arn] = last
			}
			c.mutex.Unlock()

			c.logger.Debugf("Flush apns endpoints", map[string]interface{}{
				"batch.start": batchStartIndex,
				"batch.end":   batchEndIndex,
			})

			batchStartIndex = batchEndIndex
		}
	}
}

func (c *Component) updaterSubscriptions() {
	subscriptions := []*sns.Subscription{}
	params := &sns.ListSubscriptionsInput{}

	err := c.GetSNS().ListSubscriptionsPages(params, func(p *sns.ListSubscriptionsOutput, lastPage bool) bool {
		subscriptions = append(subscriptions, p.Subscriptions...)
		return !lastPage
	})

	if err != nil {
		c.logger.Errorf("Update subscriptions error %s", err.Error())
		return
	}

	if metricSubscriptionsTotal != nil {
		metricSubscriptionsTotal.Set(float64(len(subscriptions)))
	}

	c.mutex.Lock()
	c.subscriptions = subscriptions
	c.mutex.Unlock()

	c.logger.Debugf("Updater found %d subscriptions", len(subscriptions))
}

func (c *Component) updaterTopics() {
	topics := []*sns.Topic{}
	params := &sns.ListTopicsInput{}

	err := c.GetSNS().ListTopicsPages(params, func(p *sns.ListTopicsOutput, lastPage bool) bool {
		topics = append(topics, p.Topics...)
		return !lastPage
	})

	if err != nil {
		c.logger.Errorf("Update topics error %s", err.Error())
		return
	}

	if metricTopicsTotal != nil {
		metricTopicsTotal.Set(float64(len(topics)))
	}

	c.mutex.Lock()
	c.topics = topics
	c.mutex.Unlock()

	c.logger.Debugf("Updater found %d topics", len(topics))
}

package internal

import (
	"time"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/kihamo/shadow-aws/components/aws"
)

const (
	endpointsBatchSize = 50
)

func (c *Component) GetSNS() *sns.SNS {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.services[ServiceSNS]; !ok {
		c.services[ServiceSNS] = sns.New(session.New(c.awsConfig))
	}

	return c.services[ServiceSNS].(*sns.SNS)
}

func (c *Component) GetApplications() []aws.SnsApplication {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	applications := make([]aws.SnsApplication, len(c.applications))

	i := 0
	for _, application := range c.applications {
		applications[i] = application
		i++
	}

	return applications
}

func (c *Component) GetSubscriptions() []*sns.Subscription {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	subscriptions := make([]*sns.Subscription, len(c.subscriptions))
	copy(subscriptions, c.subscriptions)

	return subscriptions
}

func (c *Component) GetTopics() []*sns.Topic {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	topics := make([]*sns.Topic, len(c.topics))
	copy(topics, c.topics)

	return topics
}

func (c *Component) loadUpdaters() {
	go func() {
		ticker := time.NewTicker(c.config.Duration(aws.ConfigUpdaterApplicationsDuration))

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
		ticker := time.NewTicker(c.config.Duration(aws.ConfigUpdaterSubscriptionsDuration))

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
		ticker := time.NewTicker(c.config.Duration(aws.ConfigUpdaterTopicsDuration))

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

	if c.config.Bool(aws.ConfigRunUpdatersOnStartup) {
		c.RunApplicationsUpdater()
		c.RunSubscriptionsUpdater()
		c.RunTopicsUpdater()
	}
}

func (c *Component) RunApplicationsUpdater() {
	c.applicationsRun <- struct{}{}
}

func (c *Component) RunSubscriptionsUpdater() {
	c.subscriptionsRun <- struct{}{}
}

func (c *Component) RunTopicsUpdater() {
	c.topicsRun <- struct{}{}
}

func (c *Component) updaterApplications() {
	lastUpdate := time.Now().UTC()
	applications := map[string]aws.SnsApplication{}
	params := &sns.ListPlatformApplicationsInput{}

	err := c.GetSNS().ListPlatformApplicationsPages(params, func(p *sns.ListPlatformApplicationsOutput, lastPage bool) bool {
		for _, a := range p.PlatformApplications {
			arn := sdk.StringValue(a.PlatformApplicationArn)

			app := aws.SnsApplication{
				Arn:                   arn,
				AwsAttributes:         make(map[string]string, len(a.Attributes)),
				Enabled:               true,
				EndpointsCount:        -1,
				EndpointsEnabledCount: -1,
				LastUpdate:            lastUpdate,
			}

			for k, v := range a.Attributes {
				app.AwsAttributes[k] = *v
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
	endpointsTotal := 0
	endpointsEnabledTotal := 0

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
			endpointsTotal += applications[i].EndpointsCount
			endpointsEnabledTotal += applications[i].EndpointsEnabledCount
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

	if metricEndpointsTotal != nil {
		metricEndpointsTotal.Set(float64(endpointsTotal))
		metricEndpointsTotal.With("status", "enabled").Set(float64(endpointsEnabledTotal))
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

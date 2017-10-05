package internal

import (
	"errors"
	"fmt"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kihamo/shadow-aws/components/aws"
)

func (c *Component) GetSES() *ses.SES {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.services[ServiceSES]; !ok {
		c.services[ServiceSES] = ses.New(session.New(c.awsConfig))
	}

	return c.services[ServiceSES].(*ses.SES)
}

func (c *Component) SendEmail(to []string, subject string, text string, html string, from string) error {
	if len(to) == 0 {
		return errors.New("To emails is empty")
	}

	if subject == "" {
		return errors.New("Subject is empty")
	}

	if text == "" && html == "" {
		return errors.New("Message is empty")
	}

	if from == "" {
		from = c.config.GetString(aws.ConfigSesFromEmail)
		name := c.config.GetString(aws.ConfigSesFromName)
		if name != "" {
			from = fmt.Sprintf("\"%s\" <%s>", name, from)
		}
	}

	input := &ses.SendEmailInput{
		Source: sdk.String(from),
		Destination: &ses.Destination{
			ToAddresses: sdk.StringSlice(to),
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: sdk.String(subject),
			},
			Body: &ses.Body{},
		},
	}

	if text != "" {
		input.Message.Body.Text = &ses.Content{
			Charset: sdk.String("UTF-8"),
			Data:    sdk.String(text),
		}
	}

	if html != "" {
		input.Message.Body.Html = &ses.Content{
			Charset: sdk.String("UTF-8"),
			Data:    sdk.String(html),
		}
	}

	_, err := c.GetSES().SendEmail(input)

	if metricSesEmailTotal != nil {
		if err != nil {
			metricSesEmailTotal.With("status", "failed").Inc()
		} else {
			metricSesEmailTotal.With("status", "success").Inc()
		}
	}

	return err
}

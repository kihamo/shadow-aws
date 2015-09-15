package service

import (
	"fmt"

	slacks "github.com/kihamo/shadow-slack/service"
	sl "github.com/nlopes/slack"
)

type AwsCommand struct {
	slacks.AbstractSlackCommand
}

func (c *AwsCommand) GetName() string {
	return "aws"
}

func (c *AwsCommand) GetDescription() string {
	return "Краткая статистика по сервису SNS в AWS"
}

func (c *AwsCommand) Run(m *sl.MessageEvent, args ...string) {
	service := c.Service.(*AwsService)
	service.mutex.RLock()

	text := fmt.Sprintf("В AWS сейчас *%d* приложений *%d* рассылок *%d* подписчиков (устойств)",
		len(service.applications),
		len(service.topics),
		len(service.subscriptions),
	)

	service.mutex.RUnlock()

	c.SendMessage(m.Channel, text)
}

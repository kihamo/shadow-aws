package service

import (
	"github.com/kihamo/shadow/service/slack"
)

func (s *AwsService) GetSlackCommands() []slack.SlackCommand {
	return []slack.SlackCommand{
		&AwsCommand{},
	}
}

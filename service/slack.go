package service

import (
	slacks "github.com/kihamo/shadow-slack/service"
)

func (s *AwsService) GetSlackCommands() []slacks.SlackCommand {
	return []slacks.SlackCommand{
		&AwsCommand{},
	}
}

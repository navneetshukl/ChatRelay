package slack

import (
	"chat-relay/internals/config"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type Slack struct {
	conf *config.Config
}

func NewSlackSvc(conf *config.Config) *Slack {
	return &Slack{
		conf: conf,
	}
}

func (s *Slack) NewSlackClient() *socketmode.Client {
	appToken := s.conf.BotConfig.SlackAppToken
	botToken := s.conf.BotConfig.SlackBotToken
	client := slack.New(botToken, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
	)
	return socketClient
}

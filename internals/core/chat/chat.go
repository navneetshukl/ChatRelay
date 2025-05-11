package chat

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type ChatRequest struct {
	UserID string
	Query  string
}

type ChatResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	Response     string `json:"full_response"`
}

type ChatUseCase interface {
	MessageEvent(client *slack.Client, event *slackevents.MessageEvent, botUserID string)
	AppMentionEvent(client *slack.Client, event *slackevents.AppMentionEvent)
}

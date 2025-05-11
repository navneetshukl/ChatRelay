package chat

import (
	"github.com/slack-go/slack/slackevents"
)

type ChatRequest struct {
	UserID string `json:"user_id"`
	Query  string `json:"query"`
	Event  string `json:"event"`
}

type ChatResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	Response     string `json:"full_response"`
}

type ChatUseCase interface {
	MessageEvent(event *slackevents.MessageEvent, botUserID string)
	AppMentionEvent(event *slackevents.AppMentionEvent,botUserID string)
}

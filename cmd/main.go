package main

import (
	"chat-relay/internals/adapter/externals/mock"
	sl "chat-relay/internals/adapter/externals/slack"
	"chat-relay/internals/config"
	"chat-relay/internals/usecase/chat"
	"log"

	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {

	// load config

	conf, err := config.LoadConfig("./.config")
	if err != nil {
		log.Panic("error in loading the config")
	}

	log.Println("Config loaded")
	slackSvc := sl.NewSlackSvc(conf)
	slackSocketClient := slackSvc.NewSlackClient()

	mockServer := mock.NewMockClientService(conf)
	chatService := chat.NewChatUseCase(conf, mockServer, &slackSocketClient.Client)

	authTest, err := slackSocketClient.Client.AuthTest()
	if err != nil {
		log.Fatalf("Error getting bot info: %v", err)
	}
	botUserID := authTest.UserID

	go func() {
		//idx := 0
		for evt := range slackSocketClient.Events {
			switch evt.Type {
			case socketmode.EventTypeInteractive:
				log.Println("Interactive event received")
			case socketmode.EventTypeEventsAPI:
				event, ok := evt.Data.(slackevents.EventsAPIEvent)
				// log.Println("##############################################")
				// log.Println("Event is ", event)
				// log.Println("Inner Event is ", event.InnerEvent)
				// log.Println("Inner Event type is ", event.InnerEvent.Type)
				// log.Println("Inner Event data is ", event.InnerEvent.Data)
				// log.Println("Event Type is ", event.Type)
				// log.Println("idx is ", idx)
				// log.Println("##############################################")
				// idx++

				if !ok {
					continue
				}
				slackSocketClient.Ack(*evt.Request)

				if event.Type == slackevents.CallbackEvent {
					switch ev := event.InnerEvent.Data.(type) {
					case *slackevents.AppMentionEvent:
						chatService.AppMentionEvent(ev,botUserID)

					case *slackevents.MessageEvent:
						if ev.SubType == "" && ev.User != botUserID {
							chatService.MessageEvent(ev, botUserID) // direct message
						}
					}
				}
			default:
				//log.Printf("Ignored event type: %s\n", evt.Type)
			}
		}
	}()

	log.Println("Slack bot is running in Socket Mode...")
	slackSocketClient.Run()
}

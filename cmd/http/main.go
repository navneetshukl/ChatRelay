package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	err := godotenv.Load("cmd/http/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	client := slack.New(botToken, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
	)

	go func() {
		idx:=0
		for evt := range socketClient.Events {
			switch evt.Type {
			case socketmode.EventTypeInteractive:
				log.Println("Interactive event received")
			case socketmode.EventTypeEventsAPI:
				event, ok := evt.Data.(slackevents.EventsAPIEvent)
				log.Println("##############################################")
				log.Println("Event is ", event)
				log.Println("Inner Event is ",event.InnerEvent)
				log.Println("Inner Event type is ",event.InnerEvent.Type)
				log.Println("Inner Event data is ",event.InnerEvent.Data)
				log.Println("Event Type is ",event.Type)
				log.Println("idx is ",idx)
				log.Println("##############################################")
				idx++

				if !ok {
					continue
				}
				socketClient.Ack(*evt.Request)

				if event.Type == slackevents.CallbackEvent {
					switch ev := event.InnerEvent.Data.(type) {
					case *slackevents.AppMentionEvent: //bot message
						AppMentionEvent(client, ev)

						//case *slackevents.MessageEvent:

						//if ev.SubType == "" { // regular message
						//	log.Printf("Message received: %s\n", ev.Text)

						//client.PostMessage(ev.Channel, slack.MsgOptionText("Simple Message: "+ev.Text, false))
						//	}
					}
				}
			default:
				//log.Printf("Ignored event type: %s\n", evt.Type)
			}
		}
	}()

	log.Println("Slack bot is running in Socket Mode...")
	socketClient.Run()
}

func AppMentionEvent(client *slack.Client, event *slackevents.AppMentionEvent) {

	// log.Println("*************************************************************************")
	// log.Println("Received Message is ", event.Text)
	// log.Println("Event Type is ", event.Type)
	// log.Println("User is ", event.User)
	// log.Println("Channel is ", event.Channel)
	// log.Println("*************************************************************************")

	// client.PostMessage(event.Channel,slack.MsgOptionText("Bot Message : Let's Play Holi 😎",false))

}

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
	// Load environment variables from .env file
	err := godotenv.Load("cmd/http/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	// Slack API and socketmode client setup
	client := slack.New(botToken, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))
	socketClient := socketmode.New(
		client,
		socketmode.OptionDebug(true),
	)

	go func() {
		for evt := range socketClient.Events {
			switch evt.Type {
			case socketmode.EventTypeInteractive:
				log.Println("Interactive event received")
			case socketmode.EventTypeEventsAPI:
				event, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					log.Printf("Ignored %+v\n", evt)
					continue
				}
				socketClient.Ack(*evt.Request)

				if event.Type == slackevents.CallbackEvent {
					switch ev := event.InnerEvent.Data.(type) {
					case *slackevents.AppMentionEvent: //bot message
						log.Printf("App mention received: %s\n", ev.Text) 
						log.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
						client.PostMessage(ev.Channel, slack.MsgOptionText("Bot Message: "+ev.Text, false))
					case *slackevents.MessageEvent:
						
						if ev.SubType == "" { // regular message
							log.Printf("Message received: %s\n", ev.Text)

							
							client.PostMessage(ev.Channel, slack.MsgOptionText("Simple Message: "+ev.Text, false))
						}
					}
				}
			default:
				log.Printf("Ignored event type: %s\n", evt.Type)
			}
		}
	}()

	log.Println("Slack bot is running in Socket Mode...")
	socketClient.Run()
}

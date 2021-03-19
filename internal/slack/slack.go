package slack

import (
	"github.com/slack-go/slack"
	"log"
)

func SendMessage(webhook, message string) {
	msg := slack.WebhookMessage{
		Text: message,
	}

	err := slack.PostWebhook(webhook, &msg)
	if err != nil {
		log.Fatalln(err)
	}
}

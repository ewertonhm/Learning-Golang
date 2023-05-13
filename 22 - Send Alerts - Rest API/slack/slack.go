package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendSlack(TextAlert string) {
	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		panic("SLACK_AUTH_TOKEN is not set")
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		panic("SLACK_CHANNEL_ID is not set")
	}

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   "danger",
		Pretext: "Alert",
		Text:    TextAlert,
	}
	_, timestampe, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestampe)
	return
}

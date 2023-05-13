package slack

import (
	"alertmanager/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

type SlackMessage struct {
	TextoAlerta string `json:"texoAlerta"`
}

func SendSlack(w http.ResponseWriter, r *http.Request) {
	var errorMessage utils.ErrorMessage

	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		log.Println("SLACK_AUTH_TOKEN is not set")
		utils.ReturnErrorMessage("SLACK_AUTH_TOKEN not defined", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Println("SLACK_CHANNEL_ID is not set")
		utils.ReturnErrorMessage("SLACK_CHANNEL_ID not defined", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	SlackMessage := SlackMessage{}
	err := json.NewDecoder(r.Body).Decode(&SlackMessage)
	if err != nil {
		log.Printf("Invalid request: %v", err)
		utils.ReturnErrorMessage("Invalid request", http.StatusBadRequest, w, &errorMessage)
		return
	}

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   "dangerhttp",
		Pretext: "Alert",
		Text:    SlackMessage.TextoAlerta,
	}
	_, timestampe, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		log.Printf("Erro ao enviar mensagem: %v", err)
		utils.ReturnErrorMessage("Erro ao enviar mensagem", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	log.Printf("Message successfully sent to channel %s at %s", channelID, timestampe)
}

/*
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
		Color:   "dangerhttp
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
*/

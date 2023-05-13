package telegram

import (
	"alertmanager/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {
	Text    string `json:"text"`
	GroupID int64  `json:"group_id"`
}

func SendTelegram(w http.ResponseWriter, r *http.Request) {

	var errorMessage utils.ErrorMessage

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not defined")
		utils.ReturnErrorMessage("TELEGRAM_BOT_TOKEN not defined", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	message := Message{}

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Printf("Invalid request: %v", err)
		utils.ReturnErrorMessage("Invalid request", http.StatusBadRequest, w, &errorMessage)
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Erro ao criar o bot: %v", err)
		utils.ReturnErrorMessage("Erro ao criar o bot", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	m := tgbotapi.NewMessage(message.GroupID, message.Text)

	retorno, err := bot.Send(m)
	if err != nil {
		log.Printf("Erro ao enviar o alerta: %v", err)
		utils.ReturnErrorMessage("Erro ao enviar o alerta", http.StatusInternalServerError, w, &errorMessage)
		return
	}
	log.Printf("Message sent: %d", retorno.MessageID)
}

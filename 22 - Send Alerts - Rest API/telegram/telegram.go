package telegram

import (
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

type ErrorMessage struct {
	Error string `json:"error"`
}

func SendTelegram(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN not defined")
	}

	var errorMessage ErrorMessage
	message := Message{}

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Printf("Invalid request: %v", err)
		errorMessage.Error = "Invalid request"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Erro ao criar o bot: %v", err)
		errorMessage.Error = "Erro ao criar o bot"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	m := tgbotapi.NewMessage(message.GroupID, message.Text)
	retorno, err := bot.Send(m)
	if err != nil {
		log.Printf("Erro ao enviar o alerta: %v", err)
		errorMessage.Error = "Erro ao enviar o alerta"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	log.Printf("Message sent: %d", retorno.MessageID)
}

/*
func SendTelegram(botApi string, message string) {
	bot, err := tgbotapi.NewBotAPI(botApi)
	if err != nil {
		panic(err)
	}

	gid := os.Getenv("TELEGRAM_GROUP_ID")
	if gid == "" {
		panic("telegram group id not defined")
	}

	integergid, err := strconv.ParseInt(gid, 10, 64)
	if err != nil {
		panic(err)
	}

	m := Message{Text: message, GroupID: integergid}
	alertText := tgbotapi.NewMessage(m.GroupID, m.Text)

	bot.Debug = true
	bot.Send(alertText)
}
*/

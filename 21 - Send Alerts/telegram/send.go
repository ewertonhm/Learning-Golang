package telegramapi

import (
	"os"
	"strconv"

	telegrambotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message struct {
	Text    string `json:"text"`
	GroupID int64  `json:"group_id"`
}

func SendTelegram(botApi string, message string) {
	bot, err := telegrambotapi.NewBotApi(botApi)
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
	alertText := telegrambotapi.NewMessage(m.GroupID, m.Text)

	bot.Debug = true
	bot.Send(alertText)
}

package main

import (
	telegram "alertmanager/telegram"
	"os"
)

func main() {
	botApi := os.Getenv("TELEGRAM_API_KEY")

	telegram.SendTelegram(botApi, "mensagem teste")
	//email.SendEmail([]string{"ewertonmarschalk@gmail.com"}, "Alerta de servidor down", "Google", "Erro ao conectar no servidor", "12/05/2023 09:24", ".\\email\\template.html")
}

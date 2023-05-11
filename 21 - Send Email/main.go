package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func main() {
	from := os.Getenv("GMAIL_EMAIL")
	password := os.Getenv("GMAIL_PASSWORD1") + os.Getenv("GMAIL_PASSWORD2")
	if len(password) < 2 || len(from) < 1 {
		fmt.Println("error: GMAIL_CREDENTIALS not set")
		os.Exit(1)
	}

	to := []string{
		from,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-versions: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Alerta: Servidor down\n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  "Google",
		Error:   "SSL Error",
		Horario: "11/05/2023 17:25:00",
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Email enviado com sucesso!")
}

package email

import (
	"alertmanager/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"text/template"
)

type Email struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Horario string   `json:"horario"`
	Server  string   `json:"server"`
	Error   string   `json:"error"`
}

func SendEmail(w http.ResponseWriter, r *http.Request) {
	var errorMessage utils.ErrorMessage

	templatePath := os.Getenv("EMAIL_TEMPLATE_PATH")
	if templatePath == "" {
		log.Printf("EMAIL_TEMPLATE_PATH not set")
		utils.ReturnErrorMessage("Internal error: EMAIL_TEMPLATE_PATH not set", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	serverSmtp := os.Getenv("EMAIL_SMTP_SERVER")
	if serverSmtp == "" {
		log.Printf("EMAIL_SMTP_SERVER not set")
		utils.ReturnErrorMessage("Internal error: EMAIL_SMTP_SERVER not set", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	portSmtp := os.Getenv("EMAIL_SMTP_PORT")
	if portSmtp == "" {
		log.Printf("EMAIL_SMTP_PORT not set")
		utils.ReturnErrorMessage("Internal error: EMAIL_SMTP_PORT not set", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	email := Email{}
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		log.Printf("Invalid request: %v", err)
		utils.ReturnErrorMessage("Invalid request", http.StatusBadRequest, w, &errorMessage)
		return
	}

	from := os.Getenv("GMAIL_EMAIL")
	if from == "" {
		log.Printf("GMAIL_EMAIL not set")
		utils.ReturnErrorMessage("Internal error: GMAIL_EMAIL not set", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		log.Printf("GMAIL_PASSWORD not set")
		utils.ReturnErrorMessage("Internal error: GMAIL_PASSWORD not set", http.StatusInternalServerError, w, &errorMessage)
		return
	}
	auth := smtp.PlainAuth("", from, os.Getenv("GMAIL_PASSWORD"), serverSmtp)

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Invalid template: %v", err)
		utils.ReturnErrorMessage("Internal error: Invalid template", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-versions: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", email.Subject, mimeHeaders)))
	t.Execute(&body, email)

	err = smtp.SendMail(fmt.Sprintf("%s:%s", serverSmtp, portSmtp), auth, from, email.To, body.Bytes())
	if err != nil {
		log.Printf("Erro ao enviar o email: %v", err)
		utils.ReturnErrorMessage("Internal error: Erro ao enviar o email", http.StatusInternalServerError, w, &errorMessage)
		return
	}
	log.Println("Email enviado com sucesso")
}

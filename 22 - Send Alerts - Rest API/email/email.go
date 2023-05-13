package email

import (
	"net/http"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {

}

/*
func SendEmail(to []string, subject string, servidor string, erro string, horario string, templatePath string) {
	from := os.Getenv("GMAIL_EMAIL")
	password := os.Getenv("GMAIL_PASSWORD1") + os.Getenv("GMAIL_PASSWORD2")
	if len(password) < 2 || len(from) < 1 {
		fmt.Println("error: GMAIL_CREDENTIALS not set")
		os.Exit(1)
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles(templatePath)

	var body bytes.Buffer

	mimeHeaders := "MIME-versions: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  servidor,
		Error:   erro,
		Horario: horario,
	})
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Email enviado com sucesso!")
}
*/

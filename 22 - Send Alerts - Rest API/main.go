package main

import (
	"alertmanager/email"
	"alertmanager/telegram"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("Starting...")

	router := mux.NewRouter()

	health := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Println(r.RemoteAddr, r.Method, r.URL, r.Proto)
		fmt.Fprintf(w, "Healthy")
	}
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/telegram", telegram.SendTelegram).Methods("POST")
	router.HandleFunc("/email", email.SendEmail).Methods("POST")
	//router.HandleFunc("/slack", slack.SendSlack).Methods("POST")
	//router.HandleFunc("/sms", sms.SendSms).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

	//router.HandleFunc("/telegram", telegram).Methods("POST")

	//botApi := os.Getenv("TELEGRAM_API_KEY")

	//telegram.SendTelegram(botApi, "mensagem teste")
	//email.SendEmail([]string{"ewertonmarschalk@gmail.com"}, "Alerta de servidor down", "Google", "Erro ao conectar no servidor", "12/05/2023 09:24", ".\\email\\template.html")
}

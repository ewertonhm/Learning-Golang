package sms

import (
	"alertmanager/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type SMSMessage struct {
	Message string `json:"message"`
	Phone   string `json:"phone"`
}

func SendSMS(w http.ResponseWriter, r *http.Request) {
	var errorMessage utils.ErrorMessage

	endpoint := os.Getenv("SMS_ENDPOINT")
	if endpoint == "" {
		log.Fatal("SMS_ENDPOINT not defined")
		utils.ReturnErrorMessage("SMS_ENDPOINT not defined", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	api_key := os.Getenv("NEXMO_API_KEY")
	if api_key == "" {
		log.Fatal("NEXMO_API_KEY not defined")
		utils.ReturnErrorMessage("NEXMO_API_KEY not defined", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	api_secret := os.Getenv("NEXMO_API_SECRET")
	if api_secret == "" {
		log.Fatal("NEXMO_API_SECRET not defined")
		utils.ReturnErrorMessage("NEXMO_API_SECRET not defined", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	message := SMSMessage{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Printf("Invalid request: %v", err)
		utils.ReturnErrorMessage("Invalid request", http.StatusBadRequest, w, &errorMessage)
		return
	}

	data := url.Values{}
	data.Set("api_key", api_key)
	data.Set("api_secret", api_secret)
	data.Set("to", message.Phone)
	data.Set("from", "Nexmo")
	data.Set("text", message.Message)

	client := &http.Client{}
	resp, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Printf("Erro ao enviar request: %v", err)
		utils.ReturnErrorMessage("Erro ao enviar request", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Printf("Erro ao enviar request: %v", err)
		utils.ReturnErrorMessage("Erro ao enviar request", http.StatusInternalServerError, w, &errorMessage)
		return
	}

	defer res.Body.Close()
	log.Println(res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erro ao enviar request: %v", err)
		utils.ReturnErrorMessage("Erro ao enviar request", http.StatusInternalServerError, w, &errorMessage)
		return
	}
	log.Println(string(body))
}

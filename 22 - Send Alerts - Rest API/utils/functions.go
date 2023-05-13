package utils

import (
	"encoding/json"
	"net/http"
)

func ReturnErrorMessage(message string, statusCode int, w http.ResponseWriter, errorMessage *ErrorMessage) {
	errorMessage.Error = message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorMessage)
}

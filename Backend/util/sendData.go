package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, StatusCode int, data interface{}) {
	w.WriteHeader(StatusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}

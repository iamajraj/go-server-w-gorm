package utils

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, msg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	b, err := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	if err != nil {
		w.Write([]byte("Something went wrong"))
	}

	w.Write(b)
}

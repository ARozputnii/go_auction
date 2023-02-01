package handlers

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if message == "record not found" {
		RespondWithJSON(w, http.StatusNotFound, map[string]string{"error": message})
		return
	}

	RespondWithJSON(w, code, map[string]string{"error": message})
}

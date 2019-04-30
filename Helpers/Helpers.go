package Helpers

import (
	"encoding/json"
	"net/http"
)

// ResponseMessage ... Type to return error message
type ResponseMessage struct {
	Message string `json:"message"`
}

// Status400 ... BadRequest
func Status400(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)

	response := ResponseMessage{Message: message}

	json.NewEncoder(w).Encode(&response)
}

// Status500 ... Server Error
func Status500(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)

	response := ResponseMessage{Message: message}

	json.NewEncoder(w).Encode(&response)
}

// Status404 ... Not Found
func Status404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(nil)
}

// Headers ... Add headers
func Headers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
}

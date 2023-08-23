package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func responseWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Printf("Responding with 5xx error %v", msg)
	}
	type errResponse struct {
		Error string `json:"error"`

	}
	reponseWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func reponseWithJSON(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to write JSON response %v", payload)
		w.WriteHeader(500)
		return 
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
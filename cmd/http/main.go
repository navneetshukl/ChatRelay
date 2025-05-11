package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BotRequest struct {
	UserID string `json:"user_id"`
	Query  string `json:"query"`
	Event  string `json:"event"`
}

type BotResponse struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message,omitempty"`
	FullResponse string `json:"full_response,omitempty"`
}

func SendChatResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(BotResponse{
			StatusCode:   http.StatusMethodNotAllowed,
			ErrorMessage: "Only POST method is allowed",
		})
		return
	}

	var req BotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BotResponse{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Invalid JSON body: " + err.Error(),
		})
		return
	}

	jsonRequest, err := json.Marshal(req)
	if err != nil {
		log.Println("Failed to marshal request:", err)
	} else {
		log.Println("Request is:", string(jsonRequest))
	}
	resp := BotResponse{
		StatusCode:   http.StatusOK,
		FullResponse: "Hi, How can I help you?",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	router := chi.NewRouter()
	router.Post("/v1/chat/stream", SendChatResponse)

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("Server failed:", err)
	}
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type BotRequest struct {
	UserID string `json:"user_id"`
	Query  string `json:"query"`
}

type BotResponse struct {
	StatusCode   int    `json:"status_code"`
	Error        string `json:"error,omitempty"`
	FullResponse string `json:"full_response,omitempty"`
}

func SendChatResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(BotResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Error:      "Only POST method is allowed",
		})
		return
	}

	var req BotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(BotResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid JSON body: " + err.Error(),
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
	http.HandleFunc("/v1/chat/stream", SendChatResponse)

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Server failed:", err)
	}
}

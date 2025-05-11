package mock

import (
	"bytes"
	"chat-relay/internals/config"
	"chat-relay/internals/core/chat"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MockClient struct {
	Config *config.Config
}

func NewMockClientService(conf *config.Config) MockClientService {
	return &MockClient{
		Config: conf,
	}
}

type MockClientService interface {
	MockServerResponse(req chat.ChatRequest) (*chat.ChatResponse, error)
}

func (srv *MockClient) MockServerResponse(req chat.ChatRequest) (*chat.ChatResponse, error) {
	url := srv.Config.MockServerConfig.BaseURL + "/v1/chat/stream"

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Println("Failed to encode JSON:", err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Request failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response:", err)
		return nil, err
	}
	log.Println("Response is ", string(body))

	var chatResp *chat.ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		log.Println("Failed to parse JSON:", err)
		return nil, err
	}
	return chatResp, nil

}

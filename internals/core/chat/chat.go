package chat

type ChatRequest struct {
	UserID string
	Query  string
}

type ChatResponse struct {
	StatusCode   int `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	Response     string `json:"full_response"`
}

package types

type jsonResponse struct {
	Error string `json:"error"`
}

type CreateMessageRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

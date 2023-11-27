package types

type CreateMessageRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type CreateProjectRequest struct {
	Title string `json:"title"`
}

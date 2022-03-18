package services

type NewMessageRequest struct {
	Uuid    string `json:"uuid"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Likes   int    `json:"likes"`
}

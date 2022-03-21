package services

type NewMessageRequest struct {
	Uuid    string `json:"uuid"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Likes   int    `json:"likes"`
}

type EditMessageRequest struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Likes   int    `json:"likes"`
}

type UpdateQuery struct {
	Uuid     string `json:"uuid"`
	Author   string `json:"author,omitempty"`
	Message  string `json:"message,omitempty"`
	Likes    int    `json:"likes,omitempty"`
	IsDelete bool   `json:"isDelete,omitempty"`
}

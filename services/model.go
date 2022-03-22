package services

type NewMessageRequest struct {
	Uuid    string `json:"uuid"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Likes   int32  `json:"likes"`
}

type EditMessageRequest struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Likes   int32  `json:"likes"`
}

type UpdateRecord struct {
	Uuid      string `json:"uuid" mapstructure:"uuid"`
	Author    string `json:"author,omitempty" mapstructure:"author,omitempty"`
	Message   string `json:"message,omitempty" mapstructure:"message,omitempty"`
	Likes     int32  `json:"likes,omitempty" mapstructure:"likes,omitempty"`
	IsDeleted bool   `json:"isDeleted,omitempty" mapstructure:"isDeleted,omitempty"`
}

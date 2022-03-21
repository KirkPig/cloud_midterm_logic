package services

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) UpdateMessageHandler(c *gin.Context) {

}

func (h *Handler) AddNewMessageHandler(c *gin.Context) {

	var req NewMessageRequest

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(409, gin.H{})
		return
	}

}

func (h *Handler) EditMessageHandler(c *gin.Context) {

	var req EditMessageRequest

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(404, gin.H{})
		return
	}

	uuid := c.Param("uuid")

}

func (h *Handler) DeleteMessageHandler(c *gin.Context) {

	uuid := c.Param("uuid")

}

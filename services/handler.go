package services

import (
	"strconv"
	"time"

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

	i, err := strconv.ParseInt(c.Param("timestamp"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{})
	}
	tm := time.Unix(i, 0)

}

func (h *Handler) AddNewMessageHandler(c *gin.Context) {

	var req NewMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(409, gin.H{
			"log": err.Error(),
		})
		return
	}

	if err := h.service.AddMessage(req); err != nil {
		c.JSON(409, gin.H{
			"log": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{})

}

func (h *Handler) EditMessageHandler(c *gin.Context) {

	var req EditMessageRequest

	req.Likes = -1

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{})
		return
	}

	uuid := c.Param("uuid")

	if err := h.service.EditMessage(uuid, req); err != nil {
		c.JSON(404, gin.H{})
		return
	}

	c.JSON(204, gin.H{})

}

func (h *Handler) DeleteMessageHandler(c *gin.Context) {

}

package services

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) UpdateMessageHandler(c *gin.Context) {

	i, err := strconv.ParseInt(c.Param("timestamp"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"log": err.Error(),
		})
		return
	}
	tm := time.Unix(i, 0)

	ups, tm, err := h.service.CheckUpdate(tm)

	if err != nil {
		c.JSON(400, gin.H{
			"log": err.Error(),
		})
		return
	}

	c.Header("Last-Sync", strconv.FormatInt(tm.Unix(), 10))
	c.JSON(200, ups)

}

func (h *Handler) AddNewMessageHandler(c *gin.Context) {

	var req NewMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(409, gin.H{
			"log": err.Error(),
		})
		return
	}

	tm, err := h.service.AddMessage(req)

	if err != nil {
		c.JSON(409, gin.H{
			"log": err.Error(),
		})
		return
	}

	c.Header("Last-Sync", strconv.FormatInt(tm.Unix(), 10))
	c.JSON(201, req)

}

func (h *Handler) EditMessageHandler(c *gin.Context) {

	var req EditMessageRequest

	req.Likes = -1

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{})
		return
	}

	uuid := c.Param("uuid")
	tm, err := h.service.EditMessage(uuid, req)

	if err != nil {
		c.JSON(404, gin.H{})
		return
	}

	c.Header("Last-Sync", strconv.FormatInt(tm.Unix(), 10))
	c.JSON(204, req)

}

func (h *Handler) DeleteMessageHandler(c *gin.Context) {

	uuid := c.Param("uuid")

	tm, err := h.service.DeleteMessage(uuid)

	if err != nil {
		c.JSON(404, gin.H{})
		return
	}

	c.Header("Last-Sync", strconv.FormatInt(tm.Unix(), 10))
	c.JSON(204, gin.H{
		"uuid": uuid,
	})

}

func (h *Handler) HealthCheck(c *gin.Context) {

	c.JSON(200, "Hello World")

}

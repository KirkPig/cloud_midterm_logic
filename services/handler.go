package services

import (
	"log"
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

func updateQueryToMap(updates []UpdateRecord) map[string]interface{} {
	updateList := make([]map[string]interface{}, 0)
	for _, up := range updates {
		updateList = append(updateList, map[string]interface{}{
			"uuid":      up.Uuid,
			"author":    map[string]interface{}{"string": up.Author},
			"message":   map[string]interface{}{"string": up.Message},
			"likes":     map[string]interface{}{"int": up.Likes},
			"isDeleted": map[string]interface{}{"boolean": up.IsDeleted},
		})
	}
	updateMap := map[string]interface{}{
		"updates": updateList,
	}
	return updateMap
}

func (h *Handler) CountMessageHandler(c *gin.Context) {
	t, err := strconv.ParseInt(c.Query("timestamp"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"log": err.Error(),
		})
		return
	}
	tm := time.Unix(t, 0)

	cnt, err := h.service.CheckUpdateCount(tm)
	c.String(200, strconv.FormatInt(cnt, 10))

}

func (h *Handler) UpdateMessageHandler(c *gin.Context) {
	log.Println("UpdateMessageHandler: Received request")
	t, err := strconv.ParseInt(c.Query("timestamp"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"log": err.Error(),
		})
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"log": err.Error(),
		})
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"log": err.Error(),
		})
		return
	}
	tm := time.Unix(t, 0)

	log.Println("UpdateMessageHandler: Checking updates", offset)
	updates, tm, err := h.service.CheckUpdate(tm, limit, offset)
	log.Println("UpdateMessageHandler: Checked updates", offset)

	if err != nil {
		c.JSON(400, gin.H{
			"log": err.Error(),
		})
		return
	}

	c.Header("Last-Sync", strconv.FormatInt(tm.Unix(), 10))
	c.JSON(200, updates)

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

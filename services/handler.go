package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/linkedin/goavro/v2"
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

	codec, err := goavro.NewCodec(`
	{
		"name": "SyncMessage",
		"namespace": "com.mycorp.mynamespace",
		"type": "record",
		"fields": [
		  {
			"name": "updates",
			"type": {
			  "type": "array",
			  "items": {
				"name": "Update",
				"type": "record",
				"fields": [
				  {
					"name": "uuid",
					"type": "string"
				  },
				  {
					"name": "author",
					"type": ["null", "string"],
					"default": null
				  },
				  {
					"name": "message",
					"type": ["null", "string"],
					"default": null
				  },
				  {
					"name": "likes",
					"type": ["null", "int"],
					"default": null
				  },
				  {
					"name": "isDeleted",
					"type": ["null", "bool"],
					"default": null
				  },
				]
			  }
			}
		  }
		]
	  }`)
	if err != nil {
		fmt.Println(err)
	}

	binary, err := codec.BinaryFromNative(nil, ups)
	if err != nil {
		fmt.Println(err)
	}

	c.Header("Last-Sync", strconv.FormatInt(tm.Unix(), 10))
	c.JSON(200, binary)

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

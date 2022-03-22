package main

import (
	"log"

	"github.com/KirkPig/cloud_midterm_logic/config"
	"github.com/KirkPig/cloud_midterm_logic/repository"
	"github.com/KirkPig/cloud_midterm_logic/services"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	conf := config.InitConfig()

	router := gin.Default()

	conn := repository.NewConnection(&conf.Postgres)
	repo := repository.NewRepository(conn)
	service := services.NewService(repo)
	apiHandler := services.NewHandler(service)
	api := router.Group("/api")
	{
		api.GET("/messages/:timestamp", apiHandler.UpdateMessageHandler)
		api.POST("/messages", apiHandler.AddNewMessageHandler)
		api.PUT("/messages/:uuid", apiHandler.EditMessageHandler)
		api.DELETE("/messages/:uuid", apiHandler.DeleteMessageHandler)
		api.GET("/health", apiHandler.HealthCheck)
	}

	log.Println("Server started on port 80")
	router.Run(":80")

}

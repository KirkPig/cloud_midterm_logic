package main

import (
	"github.com/KirkPig/cloud_midterm_logic/config"
	"github.com/KirkPig/cloud_midterm_logic/repository"
	"github.com/KirkPig/cloud_midterm_logic/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	conf := config.InitConfig()
	db := repository.New(&conf.Postgres)
	_ = repository.NewRepository(db)

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOriginFunc = func(origin string) bool { return true }
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}

	router.Use(cors.New(corsConfig))

	apiHandler := services.NewHandler(*services.NewService(*repository.NewRepository(repository.New(&config.InitConfig().Postgres))))
	api := router.Group("/api")
	{
		api.GET("/messages", apiHandler.UpdateMessageHandler)
		api.POST("/messages", apiHandler.AddNewMessageHandler)
		api.PUT("/messages/:uuid", apiHandler.EditMessageHandler)
		api.DELETE("/messages/:uuid")
	}

	router.Run(":1323")

}

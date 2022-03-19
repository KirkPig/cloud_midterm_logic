package main

import (
	"github.com/KirkPig/cloud_midterm_logic/config"
	"github.com/KirkPig/cloud_midterm_logic/repository"
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

	router.Run(":1323")

}

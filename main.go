package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewSQLConn() *gorm.DB {

	str_conn := fmt.Sprintf("host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	conn, err := gorm.Open("postgres", str_conn)

	if err != nil {
		log.Println("connection error")
		log.Fatalln(err.Error())
	}

	log.Println("db connected!! 🎉")

	return conn

}

func main() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOriginFunc = func(origin string) bool { return true }
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}

	router.Use(cors.New(config))

	router.Run(":1323")

}

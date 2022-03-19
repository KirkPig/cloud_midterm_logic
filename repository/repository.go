package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	sess *gorm.DB
}

func NewRepository(s *gorm.DB) *Repository {
	return &Repository{
		sess: s,
	}
}


type Config struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

func New(config *Config) *gorm.DB {

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.Dbname, config.Password)
	conn, err := gorm.Open("postgres", connStr)

	if err != nil {
		log.Fatalln("connection error:", err.Error())
	}

	log.Println("db connected!! 🎉")
	conn.AutoMigrate(&Message{})
	log.Println("auto migrate enabled!! 🎉")

	return conn

}
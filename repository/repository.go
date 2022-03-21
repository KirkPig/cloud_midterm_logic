package repository

import (
	"fmt"
	"log"
	"time"

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

func (r *Repository) NewMessage(uuid, author, message string, likes int, tm time.Time) error {

	return r.sess.Create(&Message{
		uuid:              uuid,
		author:            author,
		message:           message,
		likes:             likes,
		lastUpdateAuthor:  &tm,
		lastUpdateMessage: &tm,
		lastUpdateLikes:   &tm,
		isDeleted:         false,
	}).Error

}

func (r *Repository) EditMessageAuthor(uuid, author, message *string, likes *int, tm time.Time) error {

	var err error

	if author != nil {
		err = r.sess.Where(&Message{
			uuid: *uuid,
		}).Update(&Message{
			author:           *author,
			lastUpdateAuthor: &tm,
		}).Error

		if err != nil {
			return err
		}
	}

	if message != nil {
		err := r.sess.Where(&Message{
			uuid: *uuid,
		}).Update(&Message{
			message:           *message,
			lastUpdateMessage: &tm,
		}).Error

		if err != nil {
			return err
		}
	}

	if likes != nil {
		err := r.sess.Where(&Message{
			uuid: *uuid,
		}).Update(&Message{
			likes:           *likes,
			lastUpdateLikes: &tm,
		}).Error

		if err != nil {
			return err
		}
	}

	return nil

}

func (r *Repository) DeleteMessage(uuid string, tm time.Time) error {

	return r.sess.Where(&Message{
		uuid: uuid,
	}).Update(&Message{
		isDeleted: true,
	}).Error

}

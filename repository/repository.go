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

	fmt.Print(uuid, author, message, tm)

	m := Message{
		Uuid:              uuid,
		Author:            author,
		Message:           message,
		Likes:             likes,
		LastUpdateAuthor:  &tm,
		LastUpdateMessage: &tm,
		LastUpdateLikes:   &tm,
		IsDeleted:         false,
		LastUpdateDelete:  &tm,
	}

	return r.sess.Model(m).Create(&m).Error

}

func (r *Repository) EditMessage(uuid, author, message *string, likes *int, tm time.Time) error {

	var err error
	var old Message

	r.sess.Model(Message{}).Where(&Message{
		Uuid: *uuid,
	}).First(&old)

	if author != nil {

		if old.Author != *author {

			err = r.sess.Model(Message{}).Where(&Message{
				Uuid: *uuid,
			}).Update(&Message{
				Author:           *author,
				LastUpdateAuthor: &tm,
			}).Error

			if err != nil {
				return err
			}

		}

	}

	if message != nil {

		if old.Message != *message {

			err := r.sess.Model(Message{}).Where(&Message{
				Uuid: *uuid,
			}).Update(&Message{
				Message:           *message,
				LastUpdateMessage: &tm,
			}).Error

			if err != nil {
				return err
			}

		}

	}

	if likes != nil {

		if old.Likes != *likes {

			err := r.sess.Model(Message{}).Where(&Message{
				Uuid: *uuid,
			}).Update(&Message{
				Likes:           *likes,
				LastUpdateLikes: &tm,
			}).Error

			if err != nil {
				return err
			}

		}

	}

	return nil

}

func (r *Repository) DeleteMessage(uuid string, tm time.Time) error {

	return r.sess.Model(Message{}).Where(&Message{
		Uuid: uuid,
	}).Update(&Message{
		IsDeleted:        true,
		LastUpdateDelete: &tm,
	}).Error

}

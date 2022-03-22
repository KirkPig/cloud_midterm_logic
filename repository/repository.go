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

func NewConnection(config *Config) *gorm.DB {

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.Dbname, config.Password)
	log.Println("connecting to db...")
	conn, err := gorm.Open("postgres", connStr)

	if err != nil {
		log.Fatalln("connection error:", err.Error())
	}

	log.Println("db connected!! 🎉")

	return conn

}

func (r *Repository) NewMessage(uuid, author, message string, likes int32, tm time.Time) error {

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

func (r *Repository) EditMessage(uuid, author, message *string, likes *int32, tm time.Time) error {

	var old Message
	updateMessage := Message{}

	r.sess.Model(Message{}).Where(&Message{
		Uuid: *uuid,
	}).First(&old)

	if author != nil {

		if old.Author != *author {

			updateMessage.Author = *author
			updateMessage.LastUpdateAuthor = &tm

		}

	}

	if message != nil {

		if old.Message != *message {

			updateMessage.Message = *message
			updateMessage.LastUpdateMessage = &tm

		}

	}

	if likes != nil {

		if old.Likes != *likes {

			updateMessage.Likes = *likes
			updateMessage.LastUpdateLikes = &tm
		}

	}

	return r.sess.Model(Message{}).Where(&Message{
		Uuid: *uuid,
	}).Update(updateMessage).Error

}

func (r *Repository) DeleteMessage(uuid string, tm time.Time) error {

	return r.sess.Model(Message{}).Where(&Message{
		Uuid: uuid,
	}).Update(&Message{
		IsDeleted:        true,
		LastUpdateDelete: &tm,
	}).Error

}

func (r *Repository) QueryUpdate(tm time.Time, limit int64, offset int64) ([]Message, error) {

	var messages []Message
	condition := "last_update_author > ? or last_update_message > ? or last_update_likes > ? or last_update_delete > ?"
	err := r.sess.Model(Message{}).Where(condition, tm, tm, tm, tm).Order("uuid ASC").Offset(offset).Limit(limit).Find(&messages).Error
	return messages, err

}

func (r *Repository) QueryUpdateCount(tm time.Time) (int64, error) {
	var c int64
	condition := "last_update_author > ? or last_update_message > ? or last_update_likes > ? or last_update_delete > ?"
	err := r.sess.Model(Message{}).Where(condition, tm, tm, tm, tm).Count(&c).Error
	return c, err
}

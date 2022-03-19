package repository

import "time"

type Message struct {
	uuid              string     `gorm:"column:uuid;primary_key"`
	author            string     `gorm:"column:author"`
	message           string     `gorm:"column:message"`
	likes             int        `gorm:"column:likes"`
	lastUpdateAuthor  *time.Time `gorm:"column:last_update_author"`
	lastUpdateMessage *time.Time `gorm:"column:last_update_message"`
	lastUpdateLikes   *time.Time `gorm:"column:last_update_likes"`
	isDeleted         bool       `gorm:"column:is_deleted"`
}

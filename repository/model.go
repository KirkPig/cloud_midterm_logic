package repository

import "time"

type Message struct {
	Uuid              string     `gorm:"column:uuid"`
	Author            string     `gorm:"column:author"`
	Message           string     `gorm:"column:message"`
	Likes             int32      `gorm:"column:likes"`
	LastUpdateAuthor  *time.Time `gorm:"column:last_update_author"`
	LastUpdateMessage *time.Time `gorm:"column:last_update_message"`
	LastUpdateLikes   *time.Time `gorm:"column:last_update_likes"`
	IsDeleted         bool       `gorm:"column:is_deleted"`
	LastUpdateDelete  *time.Time `gorm:"column:last_update_delete"`
}

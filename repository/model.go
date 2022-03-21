package repository

import "time"

type Message struct {
	Uuid              string     `gorm:"column:uuid"`
	Author            string     `gorm:"column:author"`
	Message           string     `gorm:"column:message"`
	Likes             int        `gorm:"column:likes"`
	LastUpdateAuthor  *time.Time `gorm:"column:lastUpdateAuthor"`
	LastUpdateMessage *time.Time `gorm:"column:lastUpdateMessage"`
	LastUpdateLikes   *time.Time `gorm:"column:lastUpdateLikes"`
	IsDeleted         bool       `gorm:"column:isDeleted"`
	LastUpdateDelete  *time.Time `gorm:"column:lastUpdateDelete"`
}

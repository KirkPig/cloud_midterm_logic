package repository

import "github.com/jinzhu/gorm"

type Repository struct {
	sess *gorm.DB
}

func NewRepository(s *gorm.DB) *Repository {
	return &Repository{
		sess: s,
	}
}

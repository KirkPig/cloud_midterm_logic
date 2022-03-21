package services

import (
	"time"

	"github.com/KirkPig/cloud_midterm_logic/repository"
)

type Service struct {
	db repository.Repository
}

func NewService(db repository.Repository) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) AddMessage(req NewMessageRequest) error {

	tm := time.Now()

	return s.db.NewMessage(req.Uuid, req.Author, req.Message, req.Likes, tm)
}

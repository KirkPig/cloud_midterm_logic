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

func (s *Service) CheckUpdate(lastTM time.Time) ([]UpdateQuery, error) {

}

func (s *Service) AddMessage(req NewMessageRequest) error {

	tm := time.Now()

	return s.db.NewMessage(req.Uuid, req.Author, req.Message, req.Likes, tm)
}

func (s *Service) EditMessage(uuid string, req EditMessageRequest) error {

	tm := time.Now()
	var author *string
	var message *string
	var likes *int

	author = nil
	message = nil
	likes = nil

	if req.Author != "" {
		author = &req.Author
	}

	if req.Message != "" {
		message = &req.Message
	}

	if req.Likes != -1 {
		likes = &req.Likes
	}

	return s.db.EditMessage(&uuid, author, message, likes, tm)

}

func (s *Service) DeleteMessage(uuid string) error {

	tm := time.Now()

	return s.db.DeleteMessage(uuid, tm)

}

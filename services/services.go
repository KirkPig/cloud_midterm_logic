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

func (s *Service) CheckUpdate(lastTM time.Time) ([]UpdateQuery, time.Time, error) {

	tm := time.Now().UTC()
	msgs, err := s.db.QueryUpdate(lastTM.UTC())

	if err != nil {
		return nil, tm, err
	}

	var up []UpdateQuery

	for _, e := range msgs {

		println(e.Author, " ", e.Message, " ", e.Likes, " ", e.IsDeleted)

		k := UpdateQuery{}
		k.Uuid = e.Uuid

		if e.IsDeleted && lastTM.UTC().Unix() < e.LastUpdateAuthor.UTC().Unix() {
			k.IsDelete = e.IsDeleted
		}

		if lastTM.UTC().Unix() < e.LastUpdateAuthor.UTC().Unix() {
			k.Author = e.Author
		}

		if lastTM.UTC().Unix() < e.LastUpdateMessage.UTC().Unix() {
			k.Message = e.Message
		}

		if lastTM.UTC().Unix() < e.LastUpdateLikes.UTC().Unix() {
			k.Likes = e.Likes
		}

		up = append(up, k)

	}

	return up, tm, nil
}

func (s *Service) AddMessage(req NewMessageRequest) (time.Time, error) {

	tm := time.Now().UTC()

	err := s.db.NewMessage(req.Uuid, req.Author, req.Message, req.Likes, tm)

	return tm, err
}

func (s *Service) EditMessage(uuid string, req EditMessageRequest) (time.Time, error) {

	tm := time.Now().UTC()
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

	return tm, s.db.EditMessage(&uuid, author, message, likes, tm)

}

func (s *Service) DeleteMessage(uuid string) (time.Time, error) {

	tm := time.Now().UTC()

	return tm, s.db.DeleteMessage(uuid, tm)

}

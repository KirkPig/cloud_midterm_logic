package services

import "github.com/KirkPig/cloud_midterm_logic/repository"

type Service struct {
	db repository.Repository
}

func NewService(db repository.Repository) *Service {
	return &Service{
		db: db,
	}
}

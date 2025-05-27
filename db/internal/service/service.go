package service

import (
	"attRest/db/model"

	repository "attRest/db/internal/repo"
)

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(task *model.Task) (*model.Task, error) {
	return s.repository.Create(task)
}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *Service) Get() ([]*model.Task, error) {
	return s.repository.Get()
}

func (s *Service) MarkTaskAsDone(id int) error {
	return s.repository.MarkTaskAsDone(id)
}

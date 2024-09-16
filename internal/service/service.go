package service

import (
	"context"

	"repo/internal/repository"
)

type Service struct {
	UserRepository repository.UserRepository
}

func NewService(rep repository.UserRepository) Service {
	return Service{
		UserRepository: rep,
	}
}

func (s Service) ServiceCreate(user repository.User) {
	s.UserRepository.Create(context.Background(), user)
}

func (s Service) ServiceGetByID(userID int64) repository.User {
	user, err := s.UserRepository.GetById(context.Background(), userID)
	if err != nil {
		return repository.User{}
	}

	return user
}

func (s Service) ServiceUpdate(user repository.User) error {
	err := s.UserRepository.Update(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) ServiceDelete(id int64) error {
	return s.UserRepository.Delete(context.Background(), id)
}

func (s Service) ServiceList() ([]repository.User, error) {
	list, err := s.UserRepository.List(context.Background(), repository.Conditions{})
	return list, err
}

package services

import (
	"project/source/domain/entity"
	"project/source/infrastructure/repositories"
)

type Service struct {
	Repository repositories.UserRepository
}

func NewService(repo repositories.UserRepository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s Service) SignUpService(user entity.User) error {
	return s.Repository.CreateUser(user)
}

func (s Service) FindUserId(user entity.User) (entity.User, error) {
	return s.Repository.FindUserID(user)
}

func (s Service) FindUserPass(user entity.User) (entity.User, error) {
	return s.Repository.FindUserPasswordRepo(user)
}

func (s Service) FindUserByIdService(user_id int) (entity.User, error) {
	return s.Repository.FindUserByID(user_id)
}

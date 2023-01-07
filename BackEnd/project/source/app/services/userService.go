package services

import (
	"project/source/domain/entity"
	"project/source/infrastructure/repositories"
)

type Service struct {
	repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Repository: *repo,
	}
}

func (s Service) SignUpService(user entity.User) error {
	return s.CreateUser(user)
}

func (s Service) FindUserId(user entity.User) entity.User {
	return s.FindUserID(user)
}

func (s Service) FindUserPass(user entity.User) entity.User {
	return s.FindUserPasswordRepo(user)
}

func (s Service) FindUserByIdService(user_id int) (entity.User, error) {
	return s.FindUserByID(user_id)
}

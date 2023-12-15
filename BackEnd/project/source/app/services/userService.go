package services

import (
	"github.com/go-redis/redis/v8"
	"project/source/domain/entity"
	"project/source/infrastructure/repositories"
)

type Service struct {
	Repository  repositories.UserRepository
	RedisClient *redis.Client // Добавляем клиент Redis
}

func NewService(repo repositories.UserRepository, redisClient *redis.Client) *Service {
	return &Service{
		Repository:  repo,
		RedisClient: redisClient, // Инициализируем клиент Redis
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

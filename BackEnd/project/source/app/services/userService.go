package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"project/source/domain/entity"
	"project/source/infrastructure/repositories"
	"time"
)

type Service struct {
	Repository  repositories.UserRepository
	RedisClient RedisClient
}

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	// Добавьте другие методы, если они используются
}

func NewService(repo repositories.UserRepository, redisClient RedisClient) *Service {
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

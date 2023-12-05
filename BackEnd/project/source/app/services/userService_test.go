package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"project/source/domain/entity"
	mock_repositories "project/source/infrastructure/repositories/mock"
	"testing"
)

func TestService_SignUpService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_repositories.NewMockUserRepository(mockCtrl)
	mockUser := entity.User{Username: "adminka", UserEmail: "test@example.com", Password: "password123"}

	// Устанавливаем ожидания на мок
	mockRepo.EXPECT().CreateUser(mockUser).Return(nil)

	// Создаем экземпляр сервиса с моком репозитория
	service := NewService(mockRepo)

	// Вызываем метод сервиса
	err := service.SignUpService(mockUser)

	// Проверяем результат
	if err != nil {
		t.Errorf("SignUpService returned an error: %v", err)
	}
}

func TestService_SignUpService_InvalidUserID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_repositories.NewMockUserRepository(mockCtrl)
	invalidUser := entity.User{UserID: -1, Username: "adminka", UserEmail: "test@example.com", Password: "password123"}

	// for fail
	//mockRepo.EXPECT().CreateUser(invalidUser).Return(nil)
	mockRepo.EXPECT().CreateUser(invalidUser).Return(fmt.Errorf("invalid user ID"))

	service := NewService(mockRepo)
	err := service.SignUpService(invalidUser)

	// Проверяем, что возвращается ошибка
	if err == nil || err.Error() != "invalid user ID" {
		t.Errorf("SignUpService expected to return 'invalid user ID' error, got %v", err)
	}
}
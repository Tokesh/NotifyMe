package controller

import (
	"testing"
)

// Тест успешной валидации
func TestValidateSignUpBodySuccess(t *testing.T) {
	body := SignUpBody{
		Username:         "ValidUser123",
		UserEmail:        "valid@example.com",
		Password:         "validPass123",
		ActivationStatus: "active",
		Status:           1,
	}

	err := ValidateSignUpBody(body)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Тест неудачной валидации (неверный формат email)
func TestValidateSignUpBodyInvalidEmail(t *testing.T) {
	body := SignUpBody{
		Username:         "ValidUser123",
		UserEmail:        "invalid-email",
		Password:         "validPass123",
		ActivationStatus: "active",
		Status:           1,
	}

	err := ValidateSignUpBody(body)
	if err == nil {
		t.Errorf("Expected error, got none")
	}
}

package service

import (
	"testing"

	"github.com/jatin510/golang-hexagonal/internal/core/dto"
	"github.com/jatin510/golang-hexagonal/internal/core/model/request"
	"github.com/jatin510/golang-hexagonal/internal/core/model/response"
	"github.com/jatin510/golang-hexagonal/internal/core/port/repository"
)

type mockUserRepository struct{}

func (m *mockUserRepository) Insert(user dto.UserDTO) error {
	// simulate a duplicate user case
	if user.UserName == "test_user" {
		return repository.DuplicateUser
	}

	// simulate successful insertion
	return nil
}

func TestUserService_SignUp_Success(t *testing.T) {
	// create a mock UserRepository for testing
	userRepo := &mockUserRepository{}

	// create the UserService using the mock UserRepository
	userService := NewUserService(userRepo)

	// Test case: Successful signup
	req := &request.SignUpRequest{
		Username: "test_abc",
		Password: "12345",
	}

	res := userService.SignUp(req)

	if !res.Status {
		t.Errorf("expected status to be true, got false")
	}

	data := res.Data.(response.SignUpDataResponse)
	if data.DisplayName == "" {
		t.Errorf("expected non-empty display name, got empty")
	}
}

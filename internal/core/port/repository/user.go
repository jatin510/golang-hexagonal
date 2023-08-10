package repository

import (
	"errors"

	"github.com/jatin510/golang-hexagonal/internal/core/dto"
)

var (
	DuplicateUser = errors.New("duplicate user")
)

type UserRepository interface {
	Insert(user dto.UserDTO) error
}

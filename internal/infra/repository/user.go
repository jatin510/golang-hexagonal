package repository

import (
	"errors"
	"strings"

	"github.com/jatin510/golang-hexagonal/internal/core/dto"
	repository "github.com/jatin510/golang-hexagonal/internal/core/port/repository"
	// "user-service/internal/core/dto"
	// "user-service/internal/core/port/repository"
)


const (
	duplicateEntryMsg = "Duplicate entry"
	numberRowInserted = 1
)

var (
	insertUserErr = errors.New("failed to insert user")
)

const (
	insertUserStatement = "INSERT INTO User_test( " +
		"`username`, " +
		"`password`, " +
		"`display_name`, " +
		"`created_at`," +
		"`updated_at`) " +
		"VALUES (?, ?, ?, ?, ?)"
)

type UserRepository struct {
	db repository.Database
}

func NewUserRepository(db repository.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) Insert(user dto.UserDTO) error {
	result, err := u.db.GetDB().Exec(
		insertUserStatement,
		user.UserName,
		user.Password,
		user.DisplayName,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		if strings.Contains(err.Error(), duplicateEntryMsg) {
			return repository.DuplicateUser
		}

		return err
	}

	numRow, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numRow != numberRowInserted {
		return insertUserErr
	}

	return nil
}
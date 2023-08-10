package service

import (
	"github.com/jatin510/golang-hexagonal/internal/core/model/request"
	"github.com/jatin510/golang-hexagonal/internal/core/model/response"
)

type UserService interface {
	SignUp(request *request.SignUpRequest) *response.Response
}
package services

import (
	"github.com/ocionejr/upvote-klever/server/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService{
	return &UserService{
		userRepository: userRepository,
	}
}
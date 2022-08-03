package services

import (
	"context"
	"fmt"

	pb "github.com/ocionejr/upvote-klever/pb"
	"github.com/ocionejr/upvote-klever/server/models"
	"github.com/ocionejr/upvote-klever/server/repositories"
	"github.com/ocionejr/upvote-klever/server/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService{
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) CreateUser(in *pb.UserRequest, ctx context.Context) (*pb.UserId, error) {
	user := models.UserRequestToUser(in)

	if err := user.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid user: %v\n", err),
		)
	}

	passwordHash, err := utils.HashPassword(user.Password);
	if  err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to hash password: %v\n", err),
		)
	}
	user.Password = passwordHash

	if err := service.userRepository.InsertUser(user, ctx); err != nil {
		return nil, err
	}

	return &pb.UserId{
		Id: user.ID.Hex(),
	}, nil
}
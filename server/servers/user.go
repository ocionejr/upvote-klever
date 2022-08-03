package server

import (
	"context"

	pb "github.com/ocionejr/upvote-klever/pb"
	"github.com/ocionejr/upvote-klever/server/services"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userService *services.UserService
}

func NewUserServer(userService *services.UserService) *UserServer {
	return &UserServer{
		userService: userService,
	}
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserId, error){
	userId, err := s.userService.CreateUser(in, ctx)
	if err != nil {
		return nil, err
	}

	return userId, nil
}
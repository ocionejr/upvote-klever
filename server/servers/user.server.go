package server

import (
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
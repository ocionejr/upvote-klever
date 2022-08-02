package common

import (
	"log"
	"net"

	pb "github.com/ocionejr/upvote-klever/pb"
	server "github.com/ocionejr/upvote-klever/server/servers"
	"google.golang.org/grpc"
)

type Server struct {
	UserServer *server.UserServer
	config *Config
}

func NewServer(config *Config, userServer *server.UserServer) *Server{
	return &Server{
		UserServer: userServer,
		config: config,
	}
}

func (s *Server) RegisterServers(grpcServer *grpc.Server) {
	pb.RegisterUserServiceServer(grpcServer, s.UserServer)
}

func (s *Server) GetListener() net.Listener {
	listener, err := net.Listen("tcp", s.config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}
	return listener
} 
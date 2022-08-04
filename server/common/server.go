package common

import (
	"log"
	"net"

	pb "github.com/ocionejr/upvote-klever/pb"
	"github.com/ocionejr/upvote-klever/server/servers"
	server "github.com/ocionejr/upvote-klever/server/servers"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
	tweetServer *servers.TweetServer
}

func NewServer(config *Config, tweetServer *server.TweetServer) *Server{
	return &Server{
		config: config,
		tweetServer: tweetServer,
	}
}

func (s *Server) RegisterServers(grpcServer *grpc.Server) {
	pb.RegisterTweetServiceServer(grpcServer, s.tweetServer)
}

func (s *Server) GetListener() net.Listener {
	listener, err := net.Listen("tcp", s.config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}
	return listener
} 
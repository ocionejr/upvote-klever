package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := CreateServer()
	grpcServer := grpc.NewServer()
	server.RegisterServers(grpcServer)
	reflection.Register(grpcServer)
	listener := server.GetListener()

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err := grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

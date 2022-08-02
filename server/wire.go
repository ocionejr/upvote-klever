//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ocionejr/upvote-klever/server/common"
	"github.com/ocionejr/upvote-klever/server/repositories"
	server "github.com/ocionejr/upvote-klever/server/servers"
	"github.com/ocionejr/upvote-klever/server/services"
)

func CreateServer() *common.Server {
	panic(wire.Build(
		common.NewConfig,
		common.ConnectToDatabase,
		repositories.NewUserRepository,
		services.NewUserService,
		server.NewUserServer,
		common.NewServer,
	))
}
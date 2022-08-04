//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ocionejr/upvote-klever/server/common"
	"github.com/ocionejr/upvote-klever/server/repositories"
	server "github.com/ocionejr/upvote-klever/server/servers"
)

func CreateServer() *common.Server {
	panic(wire.Build(
		common.NewConfig,
		common.ConnectToDatabase,
		repositories.NewTweetRepository,
		server.NewTweetServer,
		common.NewServer,
	))
}

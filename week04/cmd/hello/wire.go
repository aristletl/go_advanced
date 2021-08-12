// +build wireinject

package main

import (
	"week04/internal/conf"
	"week04/internal/data"
	"week04/internal/server"
	"week04/internal/service"

	"github.com/google/wire"
)

func initApp(*conf.Data) *service.Server {
	panic(wire.Build(server.NewServer, data.NewDBModel, service.New))
}

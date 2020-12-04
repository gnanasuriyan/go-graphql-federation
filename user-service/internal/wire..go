//+build wireinject

package internal

import (
	"github.com/gnanasuriyan/go-graphql/user-micro-service/graph"
	"github.com/gnanasuriyan/go-graphql/user-micro-service/services"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(services.NewUserService)

type App struct {
	Resolver *graph.Resolver
}

func GetApp() (*App, func(), error) {
	wire.Build(
		serviceSet,
		wire.Struct(new(graph.Resolver), "*"),
		wire.Struct(new(App), "*"))
	return &App{}, nil, nil
}

//+build wireinject

package internal

import (
	"github.com/gnanasuriyan/go-graphql/product-micro-service/graph"
	"github.com/gnanasuriyan/go-graphql/product-micro-service/services"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(services.NewProductService)

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

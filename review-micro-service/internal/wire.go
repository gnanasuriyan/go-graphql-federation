//+build wireinject

package internal

import (
	"github.com/gnanasuriyan/go-graphql/review-micro-service/graph"
	"github.com/gnanasuriyan/go-graphql/review-micro-service/resolvers"
	"github.com/gnanasuriyan/go-graphql/review-micro-service/services"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(services.NewReviewService)
var resolverSet = wire.NewSet(resolvers.NewUserResolver, resolvers.NewProductResolver, resolvers.NewReviewResolver)

type App struct {
	Resolver *graph.Resolver
}

func GetApp() (*App, func(), error) {
	wire.Build(
		serviceSet,
		resolverSet,
		wire.Struct(new(graph.Resolver), "*"),
		wire.Struct(new(App), "*"))
	return &App{}, nil, nil
}

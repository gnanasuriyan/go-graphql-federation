package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/gnanasuriyan/go-graphql/user-micro-service/graph/generated"
	"github.com/gnanasuriyan/go-graphql/user-micro-service/services"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	services.IUserService
}

// Mutation returns generated.MutationResolver implementation.
type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
type queryResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Entity returns generated.EntityResolver implementation.
type entityResolver struct{ *Resolver }

func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

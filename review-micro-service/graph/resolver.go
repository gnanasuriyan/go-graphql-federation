package graph

import (
	"github.com/gnanasuriyan/go-graphql/review-micro-service/graph/generated"
	"github.com/gnanasuriyan/go-graphql/review-micro-service/resolvers"
	"github.com/gnanasuriyan/go-graphql/review-micro-service/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	services.IReviewService
	resolvers.IUserResolver
	resolvers.IProductResolver
	resolvers.IReviewResolver
}

//// Mutation returns generated.MutationResolver implementation.
//func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
//
//// Query returns generated.QueryResolver implementation.
//func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
//
//type mutationResolver struct{ *Resolver }
//type queryResolver struct{ *Resolver }

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

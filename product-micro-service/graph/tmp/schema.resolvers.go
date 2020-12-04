package tmp

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gnanasuriyan/go-graphql/product-micro-service/graph/generated"
	"github.com/gnanasuriyan/go-graphql/product-micro-service/graph/model"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, name string, price float64) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TopProducts(ctx context.Context, first *int) (*model.ProductList, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

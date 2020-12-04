package tmp

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gnanasuriyan/go-graphql/product-micro-service/graph/generated"
	"github.com/gnanasuriyan/go-graphql/product-micro-service/graph/model"
)

func (r *entityResolver) FindProductByID(ctx context.Context, id int) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

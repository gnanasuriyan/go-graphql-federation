package resolvers

import (
	"context"

	"github.com/gnanasuriyan/go-graphql/review-micro-service/graph/model"
	"github.com/google/wire"
)

type IProductResolver interface {
	FindProductByID(ctx context.Context, id int) (*model.Product, error)
}

type ProductResolver struct {
}

var NewProductResolver = wire.NewSet(wire.Struct(new(ProductResolver), "*"), wire.Bind(new(IProductResolver), new(*ProductResolver)))

func (pr *ProductResolver) FindProductByID(ctx context.Context, id int) (*model.Product, error) {
	panic("implement me")
}

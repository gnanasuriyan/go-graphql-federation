package services

import (
	"context"

	"github.com/gnanasuriyan/go-graphql/product-micro-service/graph/model"
	"github.com/google/wire"
)

type IProductService interface {
	FindProductByID(ctx context.Context, IDProduct int) (*model.Product, error)
	TopProducts(ctx context.Context, first *int) (*model.ProductList, error)
	CreateProduct(ctx context.Context, name string, price float64) (*model.Product, error)
}

type ProductService struct {
}

var productList []model.Product

var NewProductService = wire.NewSet(wire.Struct(new(ProductService), "*"), wire.Bind(new(IProductService), new(*ProductService)))

func (ps *ProductService) FindProductByID(ctx context.Context, IDProduct int) (*model.Product, error) {
	for _, product := range productList {
		if product.ID == IDProduct {
			return &product, nil
		}
	}
	return nil, nil
}

func (ps *ProductService) TopProducts(ctx context.Context, first *int) (*model.ProductList, error) {
	// TODO: implement limit
	response := &model.ProductList{
		Total: len(productList),
		Data:  []*model.Product{},
	}

	for _, productEntity := range productList {
		product := productEntity
		response.Data = append(response.Data, &product)
	}

	return response, nil
}

func (ps *ProductService) CreateProduct(ctx context.Context, name string, price float64) (*model.Product, error) {
	product := model.Product{
		ID:    len(productList) + 1,
		Name:  name,
		Price: price,
	}
	productList = append(productList, product)
	return &product, nil
}

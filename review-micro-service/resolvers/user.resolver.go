package resolvers

import (
	"context"

	"github.com/gnanasuriyan/go-graphql/review-micro-service/graph/model"
	"github.com/google/wire"
)

type IUserResolver interface {
	FindUserByID(ctx context.Context, id int) (*model.User, error)
}

type UserResolver struct {
}

var NewUserResolver = wire.NewSet(wire.Struct(new(UserResolver), "*"), wire.Bind(new(IUserResolver), new(*UserResolver)))

func (ur *UserResolver) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	panic("implement me")
}

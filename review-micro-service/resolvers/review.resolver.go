package resolvers

import (
	"context"

	"github.com/gnanasuriyan/go-graphql/review-micro-service/graph/model"
	"github.com/google/wire"
)

type IReviewResolver interface {
	FindReviewByID(ctx context.Context, id int) (*model.Review, error)
}

type ReviewResolver struct {
}

var NewReviewResolver = wire.NewSet(wire.Struct(new(ReviewResolver), "*"), wire.Bind(new(IReviewResolver), new(*ReviewResolver)))

func (rr *ReviewResolver) FindReviewByID(ctx context.Context, id int) (*model.Review, error) {
	panic("implement me")
}

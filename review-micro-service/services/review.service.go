package services

import "github.com/google/wire"

type IReviewService interface {
}

type ReviewService struct {
}

var NewReviewService = wire.NewSet(wire.Struct(new(ReviewService), "*"), wire.Bind(new(IReviewService), new(*ReviewService)))

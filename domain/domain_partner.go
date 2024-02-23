package domain

import (
	"app/model"
	"context"
)

type RepositoryInterest interface {
	Create(ctx context.Context, interest model.Interest) (model.Interest, error)
	SearchPartner(filter model.FilterInterest) ([]model.User, int, error)
}

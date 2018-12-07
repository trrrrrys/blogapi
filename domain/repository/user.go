package repository

import (
	"context"

	"blog-api/domain/model"
)

type UserRepository interface {
	GetUser(ctx context.Context) (*model.User, error)
}

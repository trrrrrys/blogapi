package infrastructure

import (
	"blog-api/domain/model"
	"blog-api/domain/repository"
	"context"
)

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (r *userRepository) GetUser(ctx context.Context) (*model.User, error) {
	return &model.User{}, nil
}

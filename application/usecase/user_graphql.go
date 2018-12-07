package usecase

import (
	"blog-api/domain/model"
	"blog-api/domain/repository"
	"context"
)

type UserUseCase interface {
	GetUser(ctx context.Context) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type userUseCase struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(cr repository.UserRepository) UserUseCase {
	return &userUseCase{
		UserRepo: cr,
	}
}

func (u *userUseCase) GetUser(ctx context.Context) (*model.User, error) {
	// 未実装
	return model.Trrrrys, nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) error {
	// 未実装
	return nil
}

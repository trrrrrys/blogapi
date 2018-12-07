package usecase

import (
	"blog-api/domain/model"
	"context"
	"fmt"
)

func (u *contentUseCase) GetContet(ctx context.Context) error {
	return nil
}
func (u *contentUseCase) GetContets(ctx context.Context, limit int) ([]*model.Content, error) {
	fmt.Println("===GetContets===")
	c, err := u.contentRepo.GetContets(ctx, limit)
	if err != nil {
		return nil, err
	}
	return c, nil
}

package usecase

import (
	"blog-api/domain/model"
	"blog-api/domain/repository"
	"context"
	"fmt"
	"time"
)

type ContentUseCase interface {
	CreateContent(context.Context, *model.Content) error
	GetContet(ctx context.Context) error
	GetContets(ctx context.Context, limit int) ([]*model.Content, error)
}

type contentUseCase struct {
	contentRepo repository.ContentRepository
}

func NewContentUsecase(cr repository.ContentRepository) ContentUseCase {
	return &contentUseCase{
		contentRepo: cr,
	}
}

func (u *contentUseCase) CreateContent(ctx context.Context, content *model.Content) error {
	fmt.Println("===CreateContent===")
	content.PublishedDate = uint(time.Now().Unix())
	return u.contentRepo.CreateContent(ctx, content)
}

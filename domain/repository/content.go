package repository

import (
	"blog-api/domain/model"
	"context"
)

type ContentRepository interface {
	CreateContent(ctx context.Context, content *model.Content) error
	GetContets(ctx context.Context, limit int) ([]*model.Content, error)
}

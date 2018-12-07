package datastore

import (
	"blog-api/domain/model"
	"blog-api/domain/repository"
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func NewContentRepository(pID string) repository.ContentRepository {
	return &contentRepository{
		projectID: pID,
	}
}

type contentRepository struct {
	projectID string
}

func (r *contentRepository) CreateContent(ctx context.Context, content *model.Content) error {
	dsClient, err := datastore.NewClient(ctx, r.projectID)
	if err != nil {
		fmt.Println(err.Error())
		return errors.Wrap(err, "datastore connection;")
	}
	tags := make([]string, len(content.Tags))
	for i, tag := range content.Tags {
		tags[i] = tag.Name
	}
	ce := &contentEntity{
		Title:         content.Title,
		Tags:          tags,
		PublishedDate: int(content.PublishedDate),
		Body:          content.Body,
	}
	if _, err := dsClient.Put(ctx, ck, ce); err != nil {
		fmt.Println(err.Error())
		return errors.Wrap(err, "datastore put;")
	}
	return nil
}

func (r *contentRepository) GetContets(ctx context.Context, limit int) ([]*model.Content, error) {
	dsClient, err := datastore.NewClient(ctx, r.projectID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.Wrap(err, "datastore connection;")
	}
	var ce []*contentEntity
	query := datastore.NewQuery(kindContent).Limit(limit)
	if _, err := dsClient.GetAll(ctx, query, &ce); err != nil {
		fmt.Println(err.Error())
		return nil, errors.Wrap(err, "datastore GetContents;")
	}

	contents := make([]*model.Content, len(ce))
	for i, c := range ce {
		tags := make([]*model.Tag, len(c.Tags))
		for index, t := range c.Tags {
			tags[index] = &model.Tag{
				Name: t,
			}
		}
		contents[i] = &model.Content{
			ID:            1,
			Title:         c.Title,
			Author:        model.Trrrrys,
			Tags:          tags,
			PublishedDate: uint(c.PublishedDate),
			Body:          c.Body,
		}
	}
	fmt.Println(contents)
	return contents, nil
}

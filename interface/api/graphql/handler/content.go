package handler

import (
	"blog-api/application/usecase"
	"errors"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type ContentHandler interface {
	Query(p graphql.ResolveParams) (interface{}, error)
}

type contentHandler struct {
	ContentUseCase usecase.ContentUseCase
}

func NewContentHandler(cu usecase.ContentUseCase) ContentHandler {
	return &contentHandler{
		ContentUseCase: cu,
	}
}

type Content struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	AuthorNick    string        `json:"author"`
	Tags          []*ContentTag `json:"tags"`
	PublishedDate uint          `json:"published_date"`
	Body          string        `json:"body"`
}

type ContentTag struct {
	Name        string `json:"tag_name"`
	Description string `json:"tag_desc"`
}

var contentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "content",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*Content)
					if ok {
						return c.ID, nil
					}
					return nil, errors.New("id error")
				},
			},
			"title": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*Content)
					if ok {
						return c.Title, nil
					}
					return nil, errors.New("title error")
				},
			},
			"author": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*Content)
					if ok {
						return c.AuthorNick, nil
					}
					return nil, errors.New("AuthorNick error")
				},
			},
			"tags": &graphql.Field{
				Type: graphql.NewList(tagType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*Content)
					if ok {
						return c.Tags, nil
					}
					return nil, errors.New("tags error")
				},
			},
			"publish_date": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*Content)
					if ok {
						return c.PublishedDate, nil
					}
					return nil, errors.New("PublishedDate error")
				},
			},
			"body": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*Content)
					if ok {
						return c.Body, nil
					}
					return nil, errors.New("body error")
				},
			},
		},
	},
)

var tagType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "tags",
		Fields: graphql.Fields{
			"tag_name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*ContentTag)
					if ok {
						return c.Name, nil
					}
					return nil, errors.New("id error")
				},
			},
			"tag_desc": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					c, ok := p.Source.(*ContentTag)
					if ok {
						return c.Description, nil
					}
					return nil, errors.New("id error")
				},
			},
		},
	})

func (h *contentHandler) Query(p graphql.ResolveParams) (interface{}, error) {
	log.Println(p)
	l, ok := p.Args["limit"].(int)
	ctx := p.Context
	if !ok {
		l = 10
	}
	contents := []*Content{}
	c, err := h.ContentUseCase.GetContets(ctx, l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, content := range c {
		var tags []*ContentTag
		for _, tag := range content.Tags {
			tags = append(tags, &ContentTag{
				Name:        tag.Name,
				Description: tag.Description,
			})
		}
		contents = append(contents, &Content{
			ID:            content.ID,
			Title:         content.Title,
			AuthorNick:    content.Author.Nick,
			PublishedDate: content.PublishedDate,
			Tags:          tags,
			Body:          content.Body,
		})
	}
	return contents, nil

}

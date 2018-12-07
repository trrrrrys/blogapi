package handler

import (
	"blog-api/application/usecase"

	"github.com/graphql-go/graphql"
)

type UserHandler interface {
	Query(p graphql.ResolveParams) (interface{}, error)
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		UserUseCase: uu,
	}
}

type User struct {
	Name        string `json:"name,omitempty"`
	Nick        string `json:"nick_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"nick_name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func (h *userHandler) Query(p graphql.ResolveParams) (interface{}, error) {
	ctx := p.Context
	u, err := h.UserUseCase.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:        u.Name,
		Nick:        u.Nick,
		Email:       u.Email,
		Description: u.Description,
	}, nil
}

// func (h *userHandler) Mutation(p graphql.ResolveParams) (interface{}, error) {
// 	ctx := p.Context
// 	user := new(model.User)
// 	user.Name, _ = p.Args["name"].(string)
// 	user.Nick, _ = p.Args["nick_name"].(string)
// 	user.Description, _ = p.Args["desc"].(string)
// 	user.Email, _ = p.Args["email"].(string)

// 	err := h.UserUseCase.UpdateUser(ctx, user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &User{
// 		Name:        user.Name,
// 		Nick:        user.Nick,
// 		Email:       user.Email,
// 		Description: user.Description,
// 	}, nil
// }

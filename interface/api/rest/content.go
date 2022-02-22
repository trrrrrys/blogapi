package rest

import (
	"blog-api/application/usecase"
	"blog-api/domain/model"
	"encoding/json"
	"log"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/unrolled/render"
)

type ContentHandler interface {
	CreateContent(w http.ResponseWriter, r *http.Request)
}

func NewContentHandler(cu usecase.ContentUseCase) ContentHandler {
	return &contentHandler{
		ContentUseCase: cu,
		rendering:      render.New(render.Options{}),
		validate:       validator.New(),
	}
}

type contentHandler struct {
	ContentUseCase usecase.ContentUseCase
	rendering      *render.Render
	validate       *validator.Validate
}

type createContetnRequest struct {
	Title string   `json:"title" validate:"required"`
	Tags  []string `json:"tags" validate:"omitempty"`
	Body  string   `json:"body" validate:"required"`
}

func (h *contentHandler) CreateContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var reqest createContetnRequest
	if err := json.NewDecoder(r.Body).Decode(&reqest); err != nil {
		log.Println(err.Error())
		h.rendering.JSON(w, http.StatusInternalServerError, &Response{err.Error()})
		return
	}
	// バリデーション
	if err := h.validate.Struct(&reqest); err != nil {
		log.Println(err.Error())
		h.rendering.JSON(w, http.StatusInternalServerError, &Response{err.Error()})
		return
	}
	tags := make([]*model.Tag, len(reqest.Tags))
	for i, t := range reqest.Tags {
		tags[i] = &model.Tag{
			Name: t,
		}
	}
	c := &model.Content{
		Title: reqest.Title,
		Tags:  tags,
		Body:  reqest.Body,
	}
	err := h.ContentUseCase.CreateContent(ctx, c)
	if err != nil {
		h.rendering.JSON(w, http.StatusInternalServerError, &Response{err.Error()})
		return
	}
	h.rendering.JSON(w, http.StatusCreated, &Response{"content created"})
}

type Response struct {
	Message string `json:"message"`
}

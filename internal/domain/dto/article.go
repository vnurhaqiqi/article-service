package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/vnurhaqiqi/go-echo-starter/shared"
)

type CreateArticleRequest struct {
	Title    string    `json:"title" validate:"required"`
	Body     string    `json:"body" validate:"required"`
	AuthorID uuid.UUID `json:"authorID" validate:"required"`
}

type ResolveArticleResponse struct {
	Articles   ArticleResponseList `json:"data"`
	Pagination struct {
		Page  int `json:"page"`
		Size  int `json:"size"`
		Total int `json:"total"`
	} `json:"pagination,omitempty"`
}

func (c CreateArticleRequest) Validate() error {
	v := shared.GetValidator()

	err := v.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type ArticleResponse struct {
	ID        uuid.UUID       `json:"id"`
	Title     string          `json:"title"`
	Body      string          `json:"body"`
	Author    *AuthorResponse `json:"author,omitempty"`
	CreatedAt time.Time       `json:"createdAt"`
}

type ArticleResponseList []ArticleResponse

func (a *ArticleResponse) SetAuthor(author AuthorResponse) {
	a.Author = &author
}

type ResolveArticleRequest struct {
	Query      null.String
	AuthorName null.String
	Page       null.String
	Size       null.String
}

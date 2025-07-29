package models

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
	"github.com/vnurhaqiqi/go-echo-starter/shared/filter"
)

type Article struct {
	ID        uuid.UUID `db:"id"`
	AuthorID  uuid.UUID `db:"author_id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`

	AuthorName string `db:"author_name"`
}

type ArticleList []Article

func (a ArticleList) GetAuthorIDs() []string {
	var authorIDs []string

	for _, article := range a {
		authorIDs = append(authorIDs, article.AuthorID.String())
	}

	return authorIDs
}

func (a ArticleList) ToResponseList() dto.ArticleResponseList {
	var resp dto.ArticleResponseList

	for _, article := range a {
		articleResponse := article.ToResponse()
		articleResponse.Author = &dto.AuthorResponse{
			ID:   article.AuthorID,
			Name: article.AuthorName,
		}

		resp = append(resp, articleResponse)
	}
	return resp
}

func (a Article) ToResponse() dto.ArticleResponse {
	return dto.ArticleResponse{
		ID:        a.ID,
		Title:     a.Title,
		Body:      a.Body,
		CreatedAt: a.CreatedAt,
		Author:    &dto.AuthorResponse{ID: a.AuthorID, Name: a.AuthorName},
	}
}

func NewArticleFromRequest(req dto.CreateArticleRequest) Article {
	return Article{
		ID:        uuid.New(),
		AuthorID:  req.AuthorID,
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
	}
}

type Paginate struct {
	Page int
	Size int
}

func (p *Paginate) SetDefaults() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Size == 0 {
		p.Size = 10
	}
}

type Sorting struct {
	SortBy        null.String `json:"field" validate:"oneof=created_at"`
	SortDirection null.String `json:"order" validate:"oneof=ASC DESC"`
}

func (s *Sorting) SetDefaults() (err error) {
	if s.SortBy.String == "" || !s.SortBy.Valid {
		s.SortBy = null.StringFrom("created_at")
	}
	if s.SortDirection.String == "" || !s.SortDirection.Valid {
		s.SortDirection = null.StringFrom("DESC")
	}
	return
}

type ArticleFilterRequest struct {
	Query      null.String
	AuthorName null.String
	Paginate   Paginate
	Sorting    Sorting
}

func (f *ArticleFilterRequest) Filter() []filter.FilterFunc {
	fn := make([]filter.FilterFunc, 0)

	if f.Query.Valid && f.Query.String != "" {
		fn = append(fn, filter.Match("(articles.title, articles.body)", f.Query.String))
	}

	if f.AuthorName.Valid && f.AuthorName.String != "" {
		fn = append(fn, filter.Like("authors.name", f.AuthorName.String))
	}

	return fn
}

func NewResolveArticleFitler(req dto.ResolveArticleRequest) ArticleFilterRequest {
	filter := ArticleFilterRequest{
		Query:      req.Query,
		AuthorName: req.AuthorName,
	}

	if req.Page.String != "" && req.Size.String != "" {
		page, _ := strconv.Atoi(req.Page.String)
		size, _ := strconv.Atoi(req.Size.String)

		filter.Paginate = Paginate{
			Page: page,
			Size: size,
		}
	}

	if req.Page.String == "" && req.Size.String == "" {
		filter.Paginate.SetDefaults()
	}
	filter.Sorting.SetDefaults()

	return filter
}

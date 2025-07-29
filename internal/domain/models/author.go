package models

import (
	"github.com/google/uuid"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
	"github.com/vnurhaqiqi/go-echo-starter/shared/filter"
)

type Author struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func NewAuthorMapByID(authors []Author) map[uuid.UUID]Author {
	authorMap := make(map[uuid.UUID]Author)

	for _, author := range authors {
		authorMap[author.ID] = author
	}
	return authorMap
}

func (a Author) ToResponse() dto.AuthorResponse {
	return dto.AuthorResponse{
		ID:   a.ID,
		Name: a.Name,
	}
}

type AuthorFilterRequest struct {
	IDs []string
}

func (f *AuthorFilterRequest) Filter() []filter.FilterFunc {
	fn := make([]filter.FilterFunc, 0)

	if len(f.IDs) > 0 {
		fn = append(fn, filter.In("id", f.IDs))
	}

	return fn
}

package dto

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateArticleRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request CreateArticleRequest
		wantErr bool
	}{
		{
			name: "empty request",
			request: CreateArticleRequest{
				Title:    "",
				Body:     "",
				AuthorID: uuid.UUID{},
			},
			wantErr: true,
		},
		{
			name: "missing title",
			request: CreateArticleRequest{
				Title:    "",
				Body:     "Test Body",
				AuthorID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "missing body",
			request: CreateArticleRequest{
				Title:    "Test Title",
				Body:     "",
				AuthorID: uuid.New(),
			},
			wantErr: true,
		},
		{
			name: "missing authorID",
			request: CreateArticleRequest{
				Title:    "Test Title",
				Body:     "Test Body",
				AuthorID: uuid.UUID{},
			},
			wantErr: true,
		},
		{
			name: "valid request",
			request: CreateArticleRequest{
				Title:    "Test Title",
				Body:     "Test Body",
				AuthorID: uuid.New(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

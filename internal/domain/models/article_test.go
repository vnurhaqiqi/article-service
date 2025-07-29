package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestArticleList_GetAuthorIDs(t *testing.T) {
	tests := []struct {
		name     string
		articles ArticleList
		want     int
	}{
		{
			name:     "empty slice",
			articles: ArticleList{},
			want:     0,
		},
		{
			name: "single article",
			articles: ArticleList{
				{
					ID:        uuid.New(),
					AuthorID:  uuid.New(),
					Title:     "Test Article",
					Body:      "Test Body",
					CreatedAt: time.Now(),
				},
			},
			want: 1,
		},
		{
			name: "multiple articles",
			articles: ArticleList{
				{
					ID:        uuid.New(),
					AuthorID:  uuid.New(),
					Title:     "Article 1",
					Body:      "Body 1",
					CreatedAt: time.Now(),
				},
				{
					ID:        uuid.New(),
					AuthorID:  uuid.New(),
					Title:     "Article 2",
					Body:      "Body 2",
					CreatedAt: time.Now(),
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.articles.GetAuthorIDs()
			assert.Equal(t, tt.want, len(result))
		})
	}
}

func TestArticleList_ToResponseList(t *testing.T) {
	tests := []struct {
		name     string
		articles ArticleList
		authors  map[uuid.UUID]Author
		want     int
	}{
		{
			name:     "empty slice",
			articles: ArticleList{},
			authors:  make(map[uuid.UUID]Author),
			want:     0,
		},
		{
			name: "single article",
			articles: ArticleList{
				{
					ID:        uuid.New(),
					AuthorID:  uuid.New(),
					Title:     "Test Article",
					Body:      "Test Body",
					CreatedAt: time.Now(),
				},
			},
			authors: make(map[uuid.UUID]Author),
			want:    1,
		},
		{
			name: "multiple articles",
			articles: ArticleList{
				{
					ID:        uuid.New(),
					AuthorID:  uuid.New(),
					Title:     "Article 1",
					Body:      "Body 1",
					CreatedAt: time.Now(),
				},
				{
					ID:        uuid.New(),
					AuthorID:  uuid.New(),
					Title:     "Article 2",
					Body:      "Body 2",
					CreatedAt: time.Now(),
				},
			},
			authors: make(map[uuid.UUID]Author),
			want:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.articles.ToResponseList()
			assert.Equal(t, tt.want, len(result))
		})
	}
}

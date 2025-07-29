package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthorMapByID(t *testing.T) {

	tests := []struct {
		name    string
		authors []Author
		want    int
	}{
		{
			name:    "empty slice",
			authors: []Author{},
			want:    0,
		},
		{
			name: "single author",
			authors: []Author{
				{
					ID:   uuid.New(),
					Name: "Test Author",
				},
			},
			want: 1,
		},
		{
			name: "multiple authors",
			authors: []Author{
				{
					ID:   uuid.New(),
					Name: "Author 1",
				},
				{
					ID:   uuid.New(),
					Name: "Author 2",
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewAuthorMapByID(tt.authors)
			assert.Equal(t, tt.want, len(result))
		})
	}
}

func TestAuthorToResponse(t *testing.T) {
	id := uuid.New()
	author := Author{
		ID:   id,
		Name: "Test Author",
	}
	response := author.ToResponse()
	assert.Equal(t, id, response.ID)
	assert.Equal(t, author.Name, response.Name)
}

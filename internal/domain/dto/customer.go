package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/vnurhaqiqi/go-echo-starter/shared"
)

type CustomerRequest struct {
	ID   uuid.UUID `json:"id" swaggerignore:"true"`
	Name string    `json:"name" validate:"required"`
}

func (r CustomerRequest) Validate() error {
	v := shared.GetValidator()

	err := v.Struct(r)
	if err != nil {
		return err
	}

	return nil
}

type CustomerResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

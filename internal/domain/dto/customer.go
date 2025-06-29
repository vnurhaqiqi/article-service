package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/models"
)

type CustomerResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCustomerResponseList(customers []models.Customer) []CustomerResponse {
	var customerResponses []CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, CustomerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		})
	}
	return customerResponses
}

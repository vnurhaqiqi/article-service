package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
)

type Customer struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewCustomerResponseList(customers []Customer) []dto.CustomerResponse {
	var customerResponses []dto.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, dto.CustomerResponse{
			ID:        customer.ID,
			Name:      customer.Name,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		})
	}
	return customerResponses
}

func NewCustomerFromRequest(req dto.CustomerRequest) Customer {
	return Customer{
		ID:        uuid.New(),
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Customer) UpdateFromRequest(req dto.CustomerRequest) {
	c.Name = req.Name
	c.UpdatedAt = time.Now()
}

func (c Customer) ToResponse() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

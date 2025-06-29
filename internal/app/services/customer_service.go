package services

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/repositories"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
)

type CustomerService interface {
	GetAllCustomer(ctx context.Context) (resp []dto.CustomerResponse, err error)
}

type CustomerServiceImpl struct {
	CustomerRepository repositories.CustomerRepository
}

func ProvideCustomerServiceImpl(customerRepository repositories.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{CustomerRepository: customerRepository}
}

func (s *CustomerServiceImpl) GetAllCustomer(ctx context.Context) (resp []dto.CustomerResponse, err error) {
	customers, err := s.CustomerRepository.FindAll(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[Customer][GetAllCustomer] error CustomerRepository.FindAll")
		return
	}

	if len(customers) == 0 {
		return
	}

	resp = dto.NewCustomerResponseList(customers)

	return
}

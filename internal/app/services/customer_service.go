package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/repositories"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/models"
)

type CustomerService interface {
	GetAllCustomer(ctx context.Context) (resp []dto.CustomerResponse, err error)
	GetCustomerByID(ctx context.Context, id uuid.UUID) (resp dto.CustomerResponse, err error)
	CreateCustomer(ctx context.Context, req dto.CustomerRequest) (err error)
	UpdateCustomer(ctx context.Context, req dto.CustomerRequest) (resp dto.CustomerResponse, err error)
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

	resp = models.NewCustomerResponseList(customers)

	return
}

func (s *CustomerServiceImpl) GetCustomerByID(ctx context.Context, id uuid.UUID) (resp dto.CustomerResponse, err error) {
	customer, err := s.CustomerRepository.FindByID(ctx, id)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[Customer][GetCustomerByID] error CustomerRepository.FindByID")
		return
	}

	resp = customer.ToResponse()

	return
}

func (s *CustomerServiceImpl) CreateCustomer(ctx context.Context, req dto.CustomerRequest) (err error) {
	customer := models.NewCustomerFromRequest(req)

	err = s.CustomerRepository.Insert(ctx, customer)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[Customer][CreateCustomer] error CustomerRepository.Insert")
		return
	}

	return
}

func (s *CustomerServiceImpl) UpdateCustomer(ctx context.Context, req dto.CustomerRequest) (resp dto.CustomerResponse, err error) {
	customer, err := s.CustomerRepository.FindByID(ctx, req.ID)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[Customer][UpdateCustomer] error CustomerRepository.FindByID")
		return
	}

	customer.UpdateFromRequest(req)

	err = s.CustomerRepository.Update(ctx, customer)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[Customer][UpdateCustomer] error CustomerRepository.Update")
		return
	}

	resp = customer.ToResponse()

	return
}

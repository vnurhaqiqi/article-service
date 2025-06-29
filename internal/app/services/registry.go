package services

import "github.com/vnurhaqiqi/go-echo-starter/internal/app/repositories"

type ServiceRegistry struct {
	CustomerService CustomerService
}

func ProvideServiceRegistry(repo *repositories.RepositoryRegistry) *ServiceRegistry {
	return &ServiceRegistry{
		CustomerService: ProvideCustomerServiceImpl(repo.CustomerRepository),
	}
}

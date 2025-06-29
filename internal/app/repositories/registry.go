package repositories

import "github.com/vnurhaqiqi/go-echo-starter/internal/infra/database"

type RepositoryRegistry struct {
	CustomerRepository CustomerRepository
}

func ProvideRepositoryRegistry(db *database.MySQLConn) *RepositoryRegistry {
	return &RepositoryRegistry{
		CustomerRepository: ProvideCustomerRepositoryImpl(*db),
	}
}

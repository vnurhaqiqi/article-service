package repositories

import (
	"context"

	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/models"
	"github.com/vnurhaqiqi/go-echo-starter/internal/infra/database"
	"github.com/vnurhaqiqi/go-echo-starter/shared/failure"
)

var (
	customerQueries = struct {
		Select string
	}{
		Select: `
		SELECT 
			id,
			name,
			created_at,
			updated_at,
		FROM 
			customer
		`,
	}
)

type CustomerRepository interface {
	FindAll(ctx context.Context) (customers []models.Customer, err error)
}

type CustomerRepositoryImpl struct {
	DB database.MySQLConn
}

func ProvideCustomerRepositoryImpl(db database.MySQLConn) CustomerRepository {
	return &CustomerRepositoryImpl{DB: db}
}

func (r *CustomerRepositoryImpl) FindAll(ctx context.Context) (customers []models.Customer, err error) {
	err = r.DB.MySQL.SelectContext(
		ctx,
		&customers,
		customerQueries.Select,
	)
	if err != nil {
		err = failure.InternalError(err)
	}

	return
}

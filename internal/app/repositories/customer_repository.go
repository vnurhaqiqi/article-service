package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/models"
	"github.com/vnurhaqiqi/go-echo-starter/internal/infra/database"
	"github.com/vnurhaqiqi/go-echo-starter/shared/failure"
)

var (
	customerQueries = struct {
		Select string
		Insert string
		Update string
	}{
		Select: `
		SELECT 
			id,
			name,
			created_at,
			updated_at
		FROM 
			customer
		`,
		Insert: `
		INSERT INTO customer 
		(
			id, 
			name, 
			created_at, 
			updated_at
		)
		VALUES 
		(
			:id, 
			:name, 
			:created_at, 
			:updated_at
		)
		`,
		Update: `
		UPDATE customer
		SET 
			name = :name, 
			updated_at = :updated_at
		WHERE 
			id = :id
		`,
	}
)

type CustomerRepository interface {
	FindAll(ctx context.Context) (customers []models.Customer, err error)
	FindByID(ctx context.Context, id uuid.UUID) (customer models.Customer, err error)
	Insert(ctx context.Context, customer models.Customer) (err error)
	Update(ctx context.Context, customer models.Customer) (err error)
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

func (r *CustomerRepositoryImpl) Insert(ctx context.Context, customer models.Customer) (err error) {
	_, err = r.DB.MySQL.NamedExecContext(
		ctx,
		customerQueries.Insert,
		customer,
	)
	if err != nil {
		err = failure.InternalError(err)
	}
	return
}

func (r *CustomerRepositoryImpl) Update(ctx context.Context, customer models.Customer) (err error) {
	_, err = r.DB.MySQL.NamedExecContext(
		ctx,
		customerQueries.Update,
		customer,
	)
	if err != nil {
		err = failure.InternalError(err)
	}
	return
}

func (r *CustomerRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (customer models.Customer, err error) {
	err = r.DB.MySQL.GetContext(
		ctx,
		&customer,
		customerQueries.Select+" WHERE id = (?)",
		id,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = failure.NotFoundFromString("customer not found")
		default:
			err = failure.InternalError(err)
		}
	}
	return
}

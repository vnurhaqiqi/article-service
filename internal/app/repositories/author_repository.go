package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/models"
	"github.com/vnurhaqiqi/go-echo-starter/internal/infra/database"
	"github.com/vnurhaqiqi/go-echo-starter/shared/failure"
	"github.com/vnurhaqiqi/go-echo-starter/shared/filter"
)

var (
	authorQueries = struct {
		Select string
	}{
		Select: `
		SELECT 
			id,
			name
		FROM 
			authors
		`,
	}
)

type AuthorRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (author models.Author, err error)
	FindByFilter(ctx context.Context, f *models.AuthorFilterRequest) (authors []models.Author, err error)
}

type AuthorRepositoryImpl struct {
	DB database.MySQLConn
}

func ProvideAuthorRepositoryImpl(db database.MySQLConn) AuthorRepository {
	return &AuthorRepositoryImpl{DB: db}
}

func (a AuthorRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (author models.Author, err error) {
	err = a.DB.MySQL.GetContext(ctx, &author, authorQueries.Select+" WHERE id = ?", id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = failure.NotFoundFromString("Author not found")
		default:
			err = failure.InternalError(err)
		}
	}
	return
}

func (a AuthorRepositoryImpl) FindByFilter(ctx context.Context, f *models.AuthorFilterRequest) (authors []models.Author, err error) {
	var (
		whereClause string
		join        string
		pagination  string
		sort        string
		args        []interface{}
	)

	if f != nil {
		f := filter.New(f.Filter()...)
		tmpArgs, clause := f.QueryClause("AND")
		if len(tmpArgs) > 0 {
			whereClause += " WHERE " + clause
		}
		args = tmpArgs

		pagination = f.Paginate()

		sort = f.SortBy()

		join = f.Joins()
	}

	query := authorQueries.Select + join + whereClause + sort + pagination

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		err = failure.InternalError(err)
		return
	}

	err = a.DB.MySQL.SelectContext(ctx, &authors, query, args...)
	if err != nil {
		err = failure.InternalError(err)
	}

	return
}

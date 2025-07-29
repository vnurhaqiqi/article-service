package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/models"
	"github.com/vnurhaqiqi/go-echo-starter/internal/infra/database"
	"github.com/vnurhaqiqi/go-echo-starter/shared/failure"
	"github.com/vnurhaqiqi/go-echo-starter/shared/filter"
)

var (
	articleQueries = struct {
		Select string
		Insert string
	}{
		Select: `
		SELECT 
			articles.id,
			articles.author_id,
			articles.title,
			articles.body,
			articles.created_at,
			authors.name as author_name
		FROM 
			articles 
		JOIN authors ON articles.author_id = authors.id
		`,
		Insert: `
		INSERT INTO articles 
		(
			id,
			title, 
			author_id,
			body, 
			created_at
		)
		VALUES 
		(
			:id,
			:title, 
			:author_id,
			:body, 
			:created_at
		)
		`,
	}
)

type ArticleRepository interface {
	Insert(ctx context.Context, article models.Article) (err error)
	FindByFilter(ctx context.Context, f *models.ArticleFilterRequest) (articles []models.Article, err error)
}

type ArticleRepositoryImpl struct {
	DB database.MySQLConn
}

func ProvideArticleRepositoryImpl(db database.MySQLConn) ArticleRepository {
	return &ArticleRepositoryImpl{DB: db}
}

func (a ArticleRepositoryImpl) Insert(ctx context.Context, article models.Article) (err error) {
	_, err = a.DB.MySQL.NamedExecContext(
		ctx,
		articleQueries.Insert,
		article,
	)
	if err != nil {
		err = failure.InternalError(err)
	}
	return
}

func (a ArticleRepositoryImpl) FindByFilter(ctx context.Context, f *models.ArticleFilterRequest) (articles []models.Article, err error) {
	var (
		whereClause string
		values      []interface{}
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
	}

	query := articleQueries.Select + whereClause + sort + pagination

	query, args, err = sqlx.In(query, values...)
	if err != nil {
		err = failure.InternalError(err)
		return
	}

	err = a.DB.MySQL.SelectContext(ctx, &articles, query, args...)
	if err != nil {
		err = failure.InternalError(err)
	}

	return
}

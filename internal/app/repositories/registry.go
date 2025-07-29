package repositories

import "github.com/vnurhaqiqi/go-echo-starter/internal/infra/database"

type RepositoryRegistry struct {
	AuthorRepository  AuthorRepository
	ArticleRepository ArticleRepository
}

func ProvideRepositoryRegistry(db *database.MySQLConn) *RepositoryRegistry {
	return &RepositoryRegistry{
		AuthorRepository:  ProvideAuthorRepositoryImpl(*db),
		ArticleRepository: ProvideArticleRepositoryImpl(*db),
	}
}

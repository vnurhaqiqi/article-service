package services

import "github.com/vnurhaqiqi/go-echo-starter/internal/app/repositories"

type ServiceRegistry struct {
	ArticleService ArticleService
}

func ProvideServiceRegistry(repo *repositories.RepositoryRegistry) *ServiceRegistry {
	return &ServiceRegistry{
		ArticleService: ProvideArticleServiceImpl(repo.ArticleRepository, repo.AuthorRepository),
	}
}

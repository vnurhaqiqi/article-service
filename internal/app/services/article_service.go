package services

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/repositories"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/models"
)

type ArticleService interface {
	Create(ctx context.Context, req dto.CreateArticleRequest) (resp dto.ArticleResponse, err error)
	ResolveByFilter(ctx context.Context, req dto.ResolveArticleRequest) (resp dto.ArticleResponseList, err error)
}

type ArticleServiceImpl struct {
	articleRepository repositories.ArticleRepository
	authorRepository  repositories.AuthorRepository
}

func ProvideArticleServiceImpl(articleRepository repositories.ArticleRepository, authorRepository repositories.AuthorRepository) ArticleService {
	return &ArticleServiceImpl{articleRepository: articleRepository, authorRepository: authorRepository}
}

func (a *ArticleServiceImpl) Create(ctx context.Context, req dto.CreateArticleRequest) (resp dto.ArticleResponse, err error) {
	// validate the author data
	author, err := a.authorRepository.FindByID(ctx, req.AuthorID)
	if err != nil {
		log.Error().
			Err(err).
			Interface("req", req).
			Msg("[ArticleService][CreateArticle] authorRepository.FindByID")
		return
	}

	article := models.NewArticleFromRequest(req)
	err = a.articleRepository.Insert(ctx, article)
	if err != nil {
		log.Error().
			Err(err).
			Interface("req", req).
			Msg("[ArticleService][CreateArticle] articleRepository.Insert")
		return
	}

	resp = article.ToResponse()
	resp.SetAuthor(author.ToResponse())

	return
}

func (a *ArticleServiceImpl) ResolveByFilter(ctx context.Context, req dto.ResolveArticleRequest) (resp dto.ArticleResponseList, err error) {
	var (
		articles models.ArticleList
		filter   = models.NewResolveArticleFitler(req)
	)

	articles, err = a.articleRepository.FindByFilter(ctx, &filter)
	if err != nil {
		log.Error().
			Err(err).
			Interface("req", req).
			Msg("[ArticleService][ResolveByFilter] error articleRepository.FindByFilter")
		return
	}

	// given empty data
	if len(articles) == 0 {
		return dto.ArticleResponseList{}, nil
	}

	resp = articles.ToResponseList()

	return
}

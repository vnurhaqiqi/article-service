package handlers

import (
	"net/http"

	"github.com/guregu/null"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/services"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
	"github.com/vnurhaqiqi/go-echo-starter/shared/failure"
	"github.com/vnurhaqiqi/go-echo-starter/shared/response"
)

type ArticleHandler struct {
	articleService services.ArticleService
}

func ProvideArticleHandler(articleService services.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) RegisterRoutes(e *echo.Group) {
	articleGroup := e.Group("/articles")
	articleGroup.POST("", h.CreateArticle)
	articleGroup.GET("", h.ResolveByFilter)
}

// @Summary Create new article
// @Description Create a new article
// @Tags articles
// @Accept json
// @Produce json
// @Param article body dto.CreateArticleRequest true "Article payload"
// @Success 201 {object} dto.ArticleResponse
// @Router /v1/articles [post]
func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	req := dto.CreateArticleRequest{}
	if err := c.Bind(&req); err != nil {
		return response.WithJSONError(failure.InternalError(err), c)
	}

	if err := req.Validate(); err != nil {
		log.Error().
			Err(err).
			Interface("req", req).
			Msg("[ArticleHandler][CreateArticle] error request validation")
		return response.WithJSONError(failure.BadRequest(err), c)
	}

	resp, err := h.articleService.Create(c.Request().Context(), req)
	if err != nil {
		return response.WithJSONError(err, c)
	}

	return response.WithJSON(http.StatusCreated, resp, c)
}

// @Summary Resolve article by filter
// @Description Resolve article by filter
// @Tags articles
// @Accept json
// @Produce json
// @Param authorName query string false "Author name"
// @Param query query string false "Query"
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Success 200 {object} dto.ArticleResponseList
// @Router /v1/articles [get]
func (h *ArticleHandler) ResolveByFilter(c echo.Context) error {
	var req dto.ResolveArticleRequest

	req.AuthorName = null.StringFrom(c.QueryParam("authorName"))
	req.Query = null.StringFrom(c.QueryParam("query"))
	req.Page = null.StringFrom(c.QueryParam("page"))
	req.Size = null.StringFrom(c.QueryParam("size"))

	resp, err := h.articleService.ResolveByFilter(c.Request().Context(), req)
	if err != nil {
		return response.WithJSONError(err, c)
	}

	return response.WithJSON(http.StatusOK, resp, c)
}

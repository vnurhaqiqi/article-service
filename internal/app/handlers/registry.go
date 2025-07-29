package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/services"
)

type HandlerRegistry struct {
	ArticleHandler *ArticleHandler
}

func ProvideHandlerRegistry(service *services.ServiceRegistry) *HandlerRegistry {
	articleHandler := ProvideArticleHandler(service.ArticleService)

	return &HandlerRegistry{
		ArticleHandler: articleHandler,
	}
}

func (h *HandlerRegistry) RegisterRoutes(e *echo.Group) {
	h.ArticleHandler.RegisterRoutes(e)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})
}

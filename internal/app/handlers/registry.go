package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/services"
)

type HandlerRegistry struct {
	CustomerHandler *CustomerHandler
}

func ProvideHandlerRegistry(service *services.ServiceRegistry) *HandlerRegistry {
	customerHandler := ProvideCustomerHandler(service.CustomerService)

	return &HandlerRegistry{
		CustomerHandler: customerHandler,
	}
}

func (h *HandlerRegistry) RegisterRoutes(e *echo.Group) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})
}

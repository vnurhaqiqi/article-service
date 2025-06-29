package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/services"
	"github.com/vnurhaqiqi/go-echo-starter/shared/response"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func ProvideCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) registerRoutes(e *echo.Group) {
	customerGroup := e.Group("/customers")
	customerGroup.GET("/", h.GetAllCustomer)
}

func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
	resp, err := h.customerService.GetAllCustomer(c.Request().Context())
	if err != nil {
		return response.WithJSONError(err, c)
	}
	return response.WithJSON(http.StatusOK, resp, c)
}

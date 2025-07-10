package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/services"
	"github.com/vnurhaqiqi/go-echo-starter/internal/domain/dto"
	"github.com/vnurhaqiqi/go-echo-starter/shared/response"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func ProvideCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) RegisterRoutes(e *echo.Group) {
	customerGroup := e.Group("/customers")
	customerGroup.GET("", h.GetAllCustomer)
	customerGroup.POST("", h.CreateCustomer)
	customerGroup.GET("/:id", h.GetCustomerByID)
	customerGroup.PUT("/:id", h.UpdateCustomer)
}

// @Summary Get all customers
// @Description Get all customers
// @Tags customers
// @Accept json
// @Produce json
// @Success 200 {object} []dto.CustomerResponse
// @Router /v1/customers [get]
func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
	resp, err := h.customerService.GetAllCustomer(c.Request().Context())
	if err != nil {
		return response.WithJSONError(err, c)
	}
	return response.WithJSON(http.StatusOK, resp, c)
}

// @Summary Get customer by ID
// @Description Get customer details by ID
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} dto.CustomerResponse
// @Router /v1/customers/{id} [get]
func (h *CustomerHandler) GetCustomerByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.WithJSONError(err, c)
	}

	resp, err := h.customerService.GetCustomerByID(c.Request().Context(), id)
	if err != nil {
		return response.WithJSONError(err, c)
	}

	return response.WithJSON(http.StatusOK, resp, c)
}

// @Summary Create new customer
// @Description Create a new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body dto.CustomerRequest true "Customer payload"
// @Success 201 {object} dto.CustomerResponse
// @Router /v1/customers [post]
func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var req dto.CustomerRequest

	if err := c.Bind(&req); err != nil {
		return response.WithJSONError(err, c)
	}

	if err := req.Validate(); err != nil {
		return response.WithJSONError(err, c)
	}

	if err := h.customerService.CreateCustomer(c.Request().Context(), req); err != nil {
		return response.WithJSONError(err, c)
	}

	return response.WithMessage(http.StatusCreated, "success", c)
}

// @Summary Update customer
// @Description Update customer details
// @Tags customers
// @Accept json
// @Produce json
// @Param id path string true "Customer ID"
// @Param customer body dto.CustomerRequest true "Customer payload"
// @Success 200 {object} dto.CustomerResponse
// @Router /v1/customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	var req dto.CustomerRequest

	if err := c.Bind(&req); err != nil {
		return response.WithJSONError(err, c)
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return response.WithJSONError(err, c)
	}
	req.ID = id

	if err := req.Validate(); err != nil {
		return response.WithJSONError(err, c)
	}

	resp, err := h.customerService.UpdateCustomer(c.Request().Context(), req)
	if err != nil {
		return response.WithJSONError(err, c)
	}

	return response.WithJSON(http.StatusOK, resp, c)
}

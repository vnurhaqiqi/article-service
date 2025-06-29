package response

import (
	"github.com/labstack/echo/v4"
	"github.com/vnurhaqiqi/go-echo-starter/shared/failure"
)

type Base struct {
	Data    *interface{} `json:"data,omitempty"`
	Error   *string      `json:"error,omitempty"`
	Message *string      `json:"message,omitempty"`
}

func WithJSON(code int, payload interface{}, c echo.Context) error {
	return c.JSON(code, Base{
		Data:    &payload,
		Message: nil,
		Error:   nil,
	})
}

func WithJSONError(err error, c echo.Context) error {
	code := failure.GetCode(err)
	message := err.Error()

	return c.JSON(code, Base{
		Error:   &message,
		Message: nil,
		Data:    nil,
	})
}

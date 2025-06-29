// wire.go
//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/vnurhaqiqi/go-echo-starter/configs"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/handlers"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/repositories"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app/services"
	"github.com/vnurhaqiqi/go-echo-starter/internal/infra/database"
)

// Initialize provides the dependencies for the application
func Initialize(cfg *configs.Config) (*echo.Echo, error) {
	wire.Build(
		// Database
		database.ProvideMySQLConn,

		wire.NewSet(
			// Repositories
			repositories.ProvideRepositoryRegistry,

			// Services
			services.ProvideServiceRegistry,

			// Handlers
			handlers.ProvideHandlerRegistry,

			// Echo
			ProvideEcho,
		),
	)
	return nil, nil
}

// ProvideEcho creates and provides a new Echo instance
func ProvideEcho(handlerRegistry *handlers.HandlerRegistry) *echo.Echo {
	e := echo.New()

	// Register health check route
	RegisterHealthCheck(e)

	// Register v1 routes
	v1 := e.Group("/v1")
	handlerRegistry.RegisterRoutes(v1)

	return e
}

// RegisterHealthCheck registers the health check route
func RegisterHealthCheck(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status": "ok",
		})
	})
}

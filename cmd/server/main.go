package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-echo-starter/configs"
	"github.com/vnurhaqiqi/go-echo-starter/internal/app"
	"github.com/vnurhaqiqi/go-echo-starter/internal/infra/logger"
)

func main() {
	logger.InitLogger()

	cfg := configs.Load()

	e, err := app.Initialize(cfg)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to initialize application")
	}

	e.Use(logger.RequestLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.Server.Cors.AllowOrigins,
		AllowMethods: cfg.Server.Cors.AllowMethods,
		AllowHeaders: cfg.Server.Cors.AllowHeaders,
	}))

	log.
		Info().
		Msg("Starting server on port " + cfg.Server.Port)

	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}

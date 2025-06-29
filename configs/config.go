package configs

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"PORT"`
		Cors struct {
			AllowOrigins []string `mapstructure:"ALLOW_ORIGINS"`
			AllowMethods []string `mapstructure:"ALLOW_METHODS"`
			AllowHeaders []string `mapstructure:"ALLOW_HEADERS"`
		} `mapstructure:"CORS"`
	} `mapstructure:"SERVER"`
	DB struct {
		MySQL struct {
			Host            string        `mapstructure:"HOST"`
			Port            string        `mapstructure:"PORT"`
			User            string        `mapstructure:"USER"`
			Password        string        `mapstructure:"PASSWORD"`
			Name            string        `mapstructure:"NAME"`
			MaxConnLifetime time.Duration `mapstructure:"MAX_CONNECTION_LIFETIME"`
			MaxIdleConn     int           `mapstructure:"MAX_IDLE_CONNECTION"`
			MaxOpenConn     int           `mapstructure:"MAX_OPEN_CONNECTION"`
			TimeZone        string        `mapstructure:"TIME_ZONE"`
		} `mapstructure:"MYSQL"`
	} `mapstructure:"DB"`
}

var (
	conf Config
	once sync.Once
)

func Load() *Config {
	once.Do(func() {
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")
		viper.SetConfigType("env")
		err := viper.ReadInConfig()

		if err != nil {
			log.Warn().
				Err(err).
				Msg("Failed reading config file, falling back to environment variables")
		}

		log.Info().
			Msg("Service configuration initialized.")

		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("")
		}
	})

	return &conf
}

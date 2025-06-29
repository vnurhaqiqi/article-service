package shared

import (
	"sync"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

var once sync.Once
var v *validator.Validate

func GetValidator() *validator.Validate {
	once.Do(func() {
		log.Info().Msg("Validator initialized.")
		v = validator.New()
	})

	return v
}

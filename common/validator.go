package common

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct validates a given struct based on the defined rules in the struct tags.
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

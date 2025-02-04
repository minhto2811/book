package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateInit() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateStruct(model interface{}) error {
	return validate.Struct(model)
}

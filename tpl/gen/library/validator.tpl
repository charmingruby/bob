package validation

import "github.com/go-playground/validator/v10"

func ValidateStructByTags(s interface{}) error {
	return validator.New().Struct(s)
}

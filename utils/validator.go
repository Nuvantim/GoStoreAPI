package utils

import (
	"github.com/go-playground/validator/v10"
)

func Validator[T any](data T) error {
	v := validator.New()
	if err := v.Struct(data); err != nil {
		return err
	}
	return nil
}

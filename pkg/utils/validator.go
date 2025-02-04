package utils

import (
	"github.com/badoux/checkmail"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func Validator[T any](data T) error {
	// validate struct
	v := validator.New()
	if err := v.Struct(data); err != nil {
		return err
	}

	val := reflect.ValueOf(data)
	emailField := val.FieldByName("Email")

	if emailField.IsValid() && emailField.String() != "" {
		if err := checkmail.ValidateFormat(emailField.String()); err != nil {
			return err
		}
	}

	return nil
}

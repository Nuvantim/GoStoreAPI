package validate

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Pre-configured validator instance
var validate = func() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := f.Tag.Get("json")
		if name == "-" {
			return ""
		}
		if idx := strings.Index(name, ","); idx != -1 {
			name = name[:idx]
		}
		if name == "" {
			return f.Name
		}
		return name
	})
	return v
}()

// Static error messages
var staticErrorMessages = map[string]string{
	"required":  "field is required",
	"email":     "invalid email format",
	"omitempty": "",
	"dive":      "",
	"default":   "invalid value",
}

// Human-readable messages for tags
func msgForTag(tag, param string) string {
	if msg, ok := staticErrorMessages[tag]; ok && msg != "" {
		return msg
	}

	switch tag {
	case "min":
		return fmt.Sprintf("minimum length is %s", param)
	case "gt":
		return fmt.Sprintf("value must be greater than %s", param)
	case "gte":
		return fmt.Sprintf("value must be ≥ %s", param)
	case "lte":
		return fmt.Sprintf("value must be ≤ %s", param)
	case "eqfield":
		return fmt.Sprintf("must be equal to %s", param)
	default:
		return staticErrorMessages["default"]
	}
}

// Validate struct
func BodyStructs[T any](data T) error {
	if err := validate.Struct(data); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) && len(ve) > 0 {
			fe := ve[0]
			return fmt.Errorf("param(%s): %s", fe.Field(), msgForTag(fe.Tag(), fe.Param()))
		}
	}
	return nil
}

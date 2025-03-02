package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

func FormatValidationError(input interface{}) map[string][]string {
	validate := validator.New()

	err := validate.RegisterValidation("password", validatePassword)
	if err != nil {
		fmt.Printf("Error registering password validator: %v\n", err)
		return nil
	}

	out := make(map[string][]string)
	if err := validate.Struct(input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			t := reflect.TypeOf(input)

			for _, fe := range validationErrors {
				fieldName := getJSONFieldName(t, fe.StructField())
				if fieldName != "" {
					out[fieldName] = append(out[fieldName], validationMessage(fe))
				}
			}
		}
	}

	return out
}

func getJSONFieldName(t reflect.Type, structField string) string {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == structField {
			return field.Tag.Get("json")
		}
	}
	return ""
}

func validationMessage(fe validator.FieldError) string {
	fieldName := FormatFieldName(fe.Field())

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s cannot be empty", fieldName)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldName)
	case "min":
		return fmt.Sprintf("%s must have at least %s characters", fieldName, fe.Param())
	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", fieldName, fe.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of the following values: %s", fieldName, fe.Param())
	case "eqfield":
		return fmt.Sprintf("%s must be the same as %s", fieldName, FormatFieldName(fe.Param()))
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", fieldName, fe.Param())
	case "numeric":
		return fmt.Sprintf("%s must only contain numbers", fieldName)
	case "datetime":
		return fmt.Sprintf("%s must be in the format DD-MM-YYYY", fieldName)
	default:
		return fe.Error()
	}
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[^\w\s]`).MatchString(password)

	return hasLower && hasUpper && hasDigit && hasSpecial
}

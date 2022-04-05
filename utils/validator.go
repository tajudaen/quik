package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func StructValidateHelper(i interface{}) (bool, []string) {
	validate := validator.New()

	err := validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var result []string
		for _, e := range errs {
			result = append(result, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}
		return true, result
	}
	return false, nil
}

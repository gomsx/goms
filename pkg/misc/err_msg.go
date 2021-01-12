package misc

import (
	e "github.com/aivuca/goms/pkg/err"

	"github.com/go-playground/validator"
)

// GetValidateError get validate error.
func GetValidateError(err error) *map[string]interface{} {
	ev := err.(validator.ValidationErrors)[0]
	field := ev.StructField()
	value := ev.Value()

	em := make(map[string]interface{})
	em["error"] = e.UserEcodeMap[field]
	em[field] = value
	return &em
}

// MapValidateError map validate error.
func MapValidateError(err error) (int64, error) {
	ev := err.(validator.ValidationErrors)[0]
	field := ev.StructField()
	return e.UserEcodeMap[field], e.UserErrMap[field]
}

// MapValidateErrorX map validate error.
func MapValidateErrorX(err error) error {
	ev := err.(validator.ValidationErrors)[0]
	field := ev.StructField()
	return e.UserErrMap[field]
}

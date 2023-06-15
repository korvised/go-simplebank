package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/korvised/go-simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is support
		return util.IsSupportedCurrency(currency)
	}

	return false
}

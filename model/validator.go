package model

import (
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"github.com/go-playground/validator/v10"
)

// NewValidator returns a new instance of 'validate' with sane defaults.
// Validate is designed to be thread-safe and used as a singleton instance.
// It caches information about your struct and validations,
// in essence only parsing your validation tags once per struct type.
// Using multiple instances neglects the benefit of caching.
//
// This alternative constructor adds a new tags to validate specific struct fields.
func NewValidator() (v *validator.Validate) {
	v = validator.New()

	if err := v.RegisterValidation("base58check", func(fl validator.FieldLevel) bool {
		if len(fl.Field().String()) == 0 {
			return false
		}
		_, _, err := base58.CheckDecode(fl.Field().String())
		return err == nil
	}); err != nil {
		panic(err)
	}

	if err := v.RegisterValidation("gt0_number", func(fl validator.FieldLevel) bool {
		v, ok := new(big.Int).SetString(fl.Field().String(), 10) //nolint:gomnd
		return ok && v.Sign() > 0
	}); err != nil {
		panic(err)
	}

	return
}

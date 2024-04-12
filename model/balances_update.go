package model

import (
	"math/big"
)

// BalancesUpdate contains a set of updated balances.
type BalancesUpdate []BalanceUpdate

// Реализация интерфейса model.Validator.

func (bu BalancesUpdate) Validate() error {
	v := NewValidator()

	for _, upd := range bu {
		if err := v.Struct(upd); err != nil {
			return err
		}
	}

	return nil
}

// -----------------------------------

// BalanceUpdate contains information about a balance update for a specific account.
type BalanceUpdate struct {
	Address    `validate:"required"`
	Account    `validate:"required"`
	Currency   `validate:"required"`
	OldValue   *big.Int `validate:"required"`
	NewValue   *big.Int `validate:"required"`
	ValueDelta *big.Int `validate:"required"`
}

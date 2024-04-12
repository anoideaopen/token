// Code generated by ifacemaker; DO NOT EDIT.

package repository

import (
	"context"
	"math/big"

	"github.com/anoideaopen/token/model"
)

// Repository describes methods, implemented by the storage package.
type Balance interface {
	// Load retrieves the balance from the database for given BalanceType, Address, and Currency.
	// If no record is found, a zero value is returned.
	Load(ctx context.Context, addr model.Address, acc model.Account, curr model.Currency) (*big.Int, error)
	// Save saves the balance to the database for given BalanceType, Address, and Currency.
	Save(ctx context.Context, addr model.Address, acc model.Account, curr model.Currency, val *big.Int) error
	// List retrieves all balances from the database for given BalanceType and Address,
	// returning them as a map where the key is the currency.
	List(ctx context.Context, addr model.Address, acc model.Account) (map[model.Currency]*big.Int, error)
}
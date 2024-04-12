// Code generated by ifacemaker; DO NOT EDIT.

package controller

import (
	"context"
	"math/big"

	"github.com/anoideaopen/token/model"
)

// Controller describes methods, implemented by the service package.
type Balance interface {
	// Deposit method is intended to increase the balance of the 'to' account.
	// The amount of increase is specified by 'val' parameter.
	Deposit(ctx context.Context, addr model.Address, acc model.Account, curr model.Currency, amt *big.Int) (bu model.BalanceUpdate, err error)
	// Withdraw method is intended to decrease the balance of the 'from' account.
	// The amount of decrease is specified by 'val' parameter.
	Withdraw(ctx context.Context, addr model.Address, acc model.Account, curr model.Currency, amt *big.Int) (bu model.BalanceUpdate, err error)
	// Transfer method is intended to move funds from one account to another.
	// The amount of funds to be moved is specified by 'val' parameter.
	Transfer(ctx context.Context, addrFrom, addrTo model.Address, acc model.Account, curr model.Currency, val *big.Int) ([2]model.BalanceUpdate, error)
	// InternalTransfer method is intended for transferring funds between two accounts
	// under the same address. The amount of funds to be moved is specified by 'val' parameter.
	InternalTransfer(ctx context.Context, addr model.Address, accFrom, accTo model.Account, curr model.Currency, val *big.Int) ([2]model.BalanceUpdate, error)
	// Fetch retrieves the balance of a specific account for a given currency.
	// It takes the context (ctx), the address (addr), the account (acc),
	// and the currency (curr) as input parameters.
	// It returns the balance as a *big.Int value and an error if something goes wrong.
	Fetch(ctx context.Context, addr model.Address, acc model.Account, curr model.Currency) (*big.Int, error)
}

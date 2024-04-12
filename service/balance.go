package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/anoideaopen/token/model"
	"github.com/anoideaopen/token/storage/repository"
)

// Balance service errors.
var (
	// ErrBalanceRepository represents a generic error related to the repository operations.
	ErrBalanceRepository = errors.New("balance repository error")

	// ErrBalanceInvalidAmount is returned when balance to store is less than zero.
	ErrBalanceInvalidAmount = errors.New("amount must be greater than 0")

	// ErrBalanceInsufficientFunds indicates insufficient funds for processing.
	ErrBalanceInsufficientFunds = errors.New("insufficient funds to process")
)

// Balance is a struct that provides methods to manipulate account balances.
//
//go:generate ifacemaker -f balance.go -o controller/balance.go -i Balance -s Balance -p controller -y "Controller describes methods, implemented by the service package."
//go:generate mockgen -package mock -source controller/balance.go -destination controller/mock/mock_balance.go
type Balance struct {
	repository.Balance
}

// Deposit method is intended to increase the balance of the 'to' account.
// The amount of increase is specified by 'val' parameter.
func (bs *Balance) Deposit(
	ctx context.Context,
	addr model.Address,
	acc model.Account,
	curr model.Currency,
	amt *big.Int,
) (bu model.BalanceUpdate, err error) {
	if amt.Sign() <= 0 {
		return bu, ErrBalanceInvalidAmount
	}

	before, err := bs.Balance.Load(ctx, addr, acc, curr)
	if err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	// balance = balance + value
	after := new(big.Int).Add(before, amt)

	if err := bs.Balance.Save(ctx, addr, acc, curr, after); err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	return model.BalanceUpdate{
		Address:    addr,
		Account:    acc,
		Currency:   curr,
		OldValue:   before,
		NewValue:   after,
		ValueDelta: amt,
	}, nil
}

// Withdraw method is intended to decrease the balance of the 'from' account.
// The amount of decrease is specified by 'val' parameter.
func (bs *Balance) Withdraw(
	ctx context.Context,
	addr model.Address,
	acc model.Account,
	curr model.Currency,
	amt *big.Int,
) (bu model.BalanceUpdate, err error) {
	if amt.Sign() <= 0 {
		return bu, ErrBalanceInvalidAmount
	}

	before, err := bs.Balance.Load(ctx, addr, acc, curr)
	if err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	// balance = balance - value
	after := new(big.Int).Sub(before, amt)

	// checking balance
	if after.Sign() < 0 {
		return bu, ErrBalanceInsufficientFunds
	}

	if err := bs.Balance.Save(ctx, addr, acc, curr, after); err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	return model.BalanceUpdate{
		Address:    addr,
		Account:    acc,
		Currency:   curr,
		OldValue:   before,
		NewValue:   after,
		ValueDelta: amt,
	}, nil
}

// Transfer method is intended to move funds from one account to another.
// The amount of funds to be moved is specified by 'val' parameter.
func (bs *Balance) Transfer(
	ctx context.Context,
	addrFrom, addrTo model.Address,
	acc model.Account,
	curr model.Currency,
	val *big.Int,
) ([2]model.BalanceUpdate, error) {
	return bs.transfer(
		ctx,
		addrFrom, addrTo,
		acc, acc,
		curr,
		val,
	)
}

// InternalTransfer method is intended for transferring funds between two accounts
// under the same address. The amount of funds to be moved is specified by 'val' parameter.
func (bs *Balance) InternalTransfer(
	ctx context.Context,
	addr model.Address,
	accFrom, accTo model.Account,
	curr model.Currency,
	val *big.Int,
) ([2]model.BalanceUpdate, error) {
	return bs.transfer(
		ctx,
		addr, addr,
		accFrom, accTo,
		curr,
		val,
	)
}

// Fetch retrieves the balance of a specific account for a given currency.
// It takes the context (ctx), the address (addr), the account (acc),
// and the currency (curr) as input parameters.
// It returns the balance as a *big.Int value and an error if something goes wrong.
func (bs *Balance) Fetch(
	ctx context.Context,
	addr model.Address,
	acc model.Account,
	curr model.Currency,
) (*big.Int, error) {
	balance, err := bs.Balance.Load(ctx, addr, acc, curr)
	if err != nil {
		return nil, bs.wrap(ErrBalanceRepository, err)
	}

	return balance, nil
}

func (bs *Balance) transfer(
	ctx context.Context,
	addrFrom, addrTo model.Address,
	accFrom, accTo model.Account,
	curr model.Currency,
	amt *big.Int,
) (bu [2]model.BalanceUpdate, err error) {
	if amt.Sign() <= 0 {
		return bu, ErrBalanceInvalidAmount
	}

	beforeFrom, err := bs.Balance.Load(ctx, addrFrom, accFrom, curr)
	if err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	beforeTo, err := bs.Balance.Load(ctx, addrTo, accTo, curr)
	if err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	// transferring [balanceFrom -> balanceTo]:
	// balanceFrom = balanceFrom - value
	// balanceTo   = balanceTo   + value
	afterFrom := new(big.Int).Sub(beforeFrom, amt)
	afterTo := new(big.Int).Add(beforeTo, amt)

	// checking balance
	if afterFrom.Sign() < 0 {
		return bu, ErrBalanceInsufficientFunds
	}

	if err := bs.Balance.Save(ctx, addrFrom, accFrom, curr, afterFrom); err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	if err := bs.Balance.Save(ctx, addrTo, accTo, curr, afterTo); err != nil {
		return bu, bs.wrap(ErrBalanceRepository, err)
	}

	return [2]model.BalanceUpdate{
		{
			Address:    addrFrom,
			Account:    accFrom,
			Currency:   curr,
			OldValue:   beforeFrom,
			NewValue:   afterFrom,
			ValueDelta: amt,
		},
		{
			Address:    addrTo,
			Account:    accTo,
			Currency:   curr,
			OldValue:   beforeTo,
			NewValue:   afterTo,
			ValueDelta: amt,
		},
	}, nil
}

func (bs *Balance) wrap(err, cause error) error {
	return fmt.Errorf("%w: %s", err, cause.Error())
}

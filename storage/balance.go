package storage

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/anoideaopen/token/model"
)

// ErrBalanceDatabase represents a generic error related to the database operations.
var ErrBalanceDatabase = errors.New("balance database error")

// Balance is a structure which encapsulates the keyvalue.DB to interact with
// balances in database.
//
//go:generate ifacemaker -f balance.go -o repository/balance.go -i Balance -s Balance -p repository -y "Repository describes methods, implemented by the storage package."
//go:generate mockgen -package mock -source repository/balance.go -destination repository/mock/mock_balance.go
type Balance struct {
	keyvalue.DB
}

// Load retrieves the balance from the database for given BalanceType, Address, and Currency.
// If no record is found, a zero value is returned.
func (b *Balance) Load(
	ctx context.Context,
	addr model.Address,
	acc model.Account,
	curr model.Currency,
) (*big.Int, error) {
	raw, err := b.DB.Get(
		ctx,
		keyvalue.Key(b.join(acc, addr, curr)),
	)
	if err != nil {
		if errors.Is(err, keyvalue.ErrNotFound) {
			return new(big.Int), nil
		}

		return nil, fmt.Errorf("%w: %s", ErrBalanceDatabase, err.Error())
	}

	if len(raw) == 0 {
		return new(big.Int), nil
	}

	return new(big.Int).SetBytes(raw), nil
}

// Save saves the balance to the database for given BalanceType, Address, and Currency.
func (b *Balance) Save(
	ctx context.Context,
	addr model.Address,
	acc model.Account,
	curr model.Currency,
	val *big.Int,
) error {
	if err := b.DB.Set(
		ctx,
		keyvalue.Key(b.join(acc, addr, curr)),
		keyvalue.Value(val.Bytes()),
	); err != nil {
		return fmt.Errorf("%w: %s", ErrBalanceDatabase, err.Error())
	}

	return nil
}

// List retrieves all balances from the database for given BalanceType and Address,
// returning them as a map where the key is the currency.
func (b *Balance) List(
	ctx context.Context,
	addr model.Address,
	acc model.Account,
) (map[model.Currency]*big.Int, error) {
	// example: "4f/address"
	iter, err := b.DB.Iter(ctx, keyvalue.Prefix(b.join(acc, addr, "")))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrBalanceDatabase, err.Error())
	}
	defer iter.Close()

	out := make(map[model.Currency]*big.Int)
	for iter.HasNext() {
		k, v, err := iter.Next()
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrBalanceDatabase, err.Error())
		}

		keys := strings.Split(string(k), keyvalue.KeySeparator)
		if len(keys) != 3 { //nolint:gomnd
			return nil, fmt.Errorf(
				"%w: invalid iterator's key '%s'",
				ErrBalanceDatabase,
				k,
			)
		}

		b := new(big.Int)
		if len(v) > 0 {
			b.SetBytes(v)
		}

		out[model.Currency(keys[2])] = b
	}

	return out, nil
}

// join creates a unique key for the database record based on the BalanceType,
// Address, and Currency.
// example: "4f/address/currency" or "4f/address"
func (b *Balance) join(
	acc model.Account,
	addr model.Address,
	curr model.Currency,
) string {
	return keyvalue.Join(
		b.hex(acc),
		string(addr),
		string(curr),
	)
}

func (b *Balance) hex(acc model.Account) string {
	return hex.EncodeToString([]byte{byte(acc)})
}

package model

import (
	"fmt"
)

// Address represents a unique identifier for a specific user.
type Address string

// Account is used to categorize different types of accounts associated with an address.
type Account int

// Constants for Account Types.
const (
	AccountToken         Account = 43
	AccountTokenLocked   Account = 46
	AccountAllowed       Account = 44
	AccountAllowedLocked Account = 47
)

// String returns a string representation of the Account.
func (a Account) String() string {
	switch a {
	case AccountToken:
		return "AccountToken"
	case AccountTokenLocked:
		return "AccountTokenLocked"
	case AccountAllowed:
		return "AccountAllowed"
	case AccountAllowedLocked:
		return "AccountAllowedLocked"
	default:
		return fmt.Sprintf("Unknown Account type: %d", a)
	}
}

// Currency represents the name of the asset used to store balances in an account.
type Currency string

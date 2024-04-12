package service

import (
	"context"
	"math/big"
	"testing"

	"github.com/anoideaopen/token/model"
	ctrl "github.com/anoideaopen/token/service/controller/mock"
	repo "github.com/anoideaopen/token/storage/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type account struct {
	account  model.Account
	currency model.Currency
	balance  *big.Int
}

type user struct {
	address  model.Address
	account1 account
	account2 account
}

var (
	ctx = context.Background()

	user1 = user{
		address: "naBqaB46uCQxNQgLbCpMrVrHS694G9iLw78LFwvsM6duEzpAK",
		account1: account{
			account:  model.AccountAllowed,
			currency: "USD",
			balance:  big.NewInt(100),
		},
		account2: account{
			account:  model.AccountAllowedLocked,
			currency: "USD",
			balance:  big.NewInt(200),
		},
	}

	user2 = user{
		address: "MtsMcupUJWWrCEWd1j3EoFKct87CrYMxrwTT3qjLP9TBZwdbk",
		account1: account{
			account:  model.AccountAllowed,
			currency: "USD",
			balance:  big.NewInt(300),
		},
		account2: account{
			account:  model.AccountAllowedLocked,
			currency: "USD",
			balance:  big.NewInt(400),
		},
	}
)

type environment struct {
	assert      *assert.Assertions
	ctrlGomock  *gomock.Controller
	repoBalance *repo.MockBalance
	ctrlBalance *ctrl.MockBalance
}

func newEnvironment(t *testing.T) *environment {
	ctrlGomock := gomock.NewController(t)
	return &environment{
		assert:      assert.New(t),
		ctrlGomock:  ctrlGomock,
		repoBalance: repo.NewMockBalance(ctrlGomock),
		ctrlBalance: ctrl.NewMockBalance(ctrlGomock),
	}
}

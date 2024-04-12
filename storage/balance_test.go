//nolint:gci
package storage

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/anoideaopen/token/keyvalue/mock"
	"github.com/anoideaopen/token/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBalance_Load(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDB(ctrl)

	b := &Balance{
		DB: mockDB,
	}

	tt := model.AccountToken
	a := model.Address("0x123")
	c := model.Currency("ETH")

	mockDB.EXPECT().Get(
		gomock.Any(),
		keyvalue.Key(b.join(tt, a, c)),
	).Return(big.NewInt(100).Bytes(), nil)

	res, err := b.Load(context.Background(), a, tt, c)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(100), res)
}

func TestBalance_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDB(ctrl)

	b := &Balance{
		DB: mockDB,
	}

	tt := model.AccountToken
	a := model.Address("0x123")
	c := model.Currency("ETH")

	mockDB.EXPECT().Set(
		gomock.Any(),
		keyvalue.Key(b.join(tt, a, c)),
		gomock.Any(),
	).Return(nil)

	err := b.Save(context.Background(), a, tt, c, big.NewInt(100))
	assert.NoError(t, err)
}

func TestBalance_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDB(ctrl)
	mockIter := mock.NewMockIterator(ctrl)

	b := &Balance{
		DB: mockDB,
	}

	tt := model.AccountToken
	a := model.Address("0x123")
	c := model.Currency("ETH")

	mockDB.EXPECT().
		Iter(
			gomock.Any(),
			keyvalue.Prefix(b.join(tt, a, "")),
		).Return(mockIter, nil)

	mockIter.EXPECT().HasNext().Return(true).Times(1)
	mockIter.EXPECT().HasNext().Return(false).Times(1)
	mockIter.EXPECT().
		Next().
		Return(keyvalue.Key(fmt.Sprintf("%x/%s/%s", tt, a, c)), big.NewInt(100).Bytes(), nil)
	mockIter.EXPECT().Close()

	res, err := b.List(context.Background(), a, tt)
	assert.NoError(t, err)
	assert.Equal(t, map[model.Currency]*big.Int{
		c: big.NewInt(100),
	}, res)
}

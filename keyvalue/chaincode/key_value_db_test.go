package chaincode

import (
	"context"
	"testing"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/anoideaopen/token/keyvalue/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestKeyValueDB_SaveLoadDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stub := mock.NewMockChaincodeStubInterface(ctrl)
	cs := &KeyValueDB{Stub: stub}
	ctx := context.Background()

	// Setting up the mock object
	stub.EXPECT().CreateCompositeKey("key1", []string{"key2"}).Return("key1\x00key2", nil).Times(3)
	stub.EXPECT().PutState("key1\x00key2", []byte("value")).Return(nil)
	stub.EXPECT().PutState("key3", []byte("value")).Return(nil)
	stub.EXPECT().GetState("key1\x00key2").Return([]byte("value"), nil)
	stub.EXPECT().GetState("key3").Return([]byte("value"), nil)
	stub.EXPECT().DelState("key1\x00key2").Return(nil)

	// Testing Save
	err := cs.Set(ctx, "key1/key2", keyvalue.Value("value"))
	assert.Nil(t, err)

	err = cs.Set(ctx, "key3", keyvalue.Value("value"))
	assert.Nil(t, err)

	// Testing Load
	value, err := cs.Get(ctx, "key1/key2")
	assert.Nil(t, err)
	assert.Equal(t, keyvalue.Value("value"), value)

	value, err = cs.Get(ctx, "key3")
	assert.Nil(t, err)
	assert.Equal(t, keyvalue.Value("value"), value)

	// Testing Delete
	err = cs.Del(ctx, "key1/key2")
	assert.Nil(t, err)
}

func TestKeyValueDB_Iterate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stub := mock.NewMockChaincodeStubInterface(ctrl)
	cs := &KeyValueDB{Stub: stub}
	ctx := context.Background()

	// Setting up the mock objects
	gomock.InOrder(
		stub.EXPECT().GetStateByPartialCompositeKey("key1", []string{}).Return(nil, nil),
		stub.EXPECT().
			GetStateByPartialCompositeKey("key1", []string{"key2", "key3"}).
			Return(nil, nil),
	)

	// Testing Iterate
	iterator, err := cs.Iter(ctx, "key1")
	assert.Nil(t, err)
	assert.NotNil(t, iterator)

	iterator, err = cs.Iter(ctx, "key1/key2/key3")
	assert.Nil(t, err)
	assert.NotNil(t, iterator)
}

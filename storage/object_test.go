package storage

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/anoideaopen/token/keyvalue/mock"
	"github.com/anoideaopen/token/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestObject_Load(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDB(ctrl)

	o := &Object{
		DB: mockDB,
	}

	key := model.ObjectQuery("testKey")

	var (
		to      = new(testObject)
		blob, _ = json.Marshal(testObject1)
	)

	mockDB.EXPECT().Get(gomock.Any(), keyvalue.Key(key)).Return(blob, nil)

	err := o.Load(context.Background(), key, to)
	assert.NoError(t, err)
	assert.Equal(t, to, testObject1)
}

func TestObject_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDB(ctrl)

	o := &Object{
		DB: mockDB,
	}

	key := model.ObjectQuery("testKey")

	blob, _ := json.Marshal(testObject1)
	mockDB.EXPECT().Set(gomock.Any(), keyvalue.Key(key), blob).Return(nil)

	err := o.Save(context.Background(), key, testObject1)
	assert.NoError(t, err)
}

func TestObject_Iter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockDB(ctrl)
	mockIter := mock.NewMockIterator(ctrl)

	o := &Object{
		DB: mockDB,
	}

	prefix := model.ObjectQuery("testPrefix")

	var (
		template = testObject1
		key      = keyvalue.Key("testKey")
		blob, _  = json.Marshal(testObject1)
	)

	mockDB.EXPECT().Iter(gomock.Any(), keyvalue.Prefix(prefix)).Return(mockIter, nil)
	mockIter.EXPECT().HasNext().Return(true).Times(1)
	mockIter.EXPECT().Next().Return(key, blob, nil).Times(1)
	mockIter.EXPECT().HasNext().Return(false).Times(1) // This is to end the loop in Iter method
	mockIter.EXPECT().Close().Return(nil).Times(1)

	count := 0
	err := o.Iter(context.Background(), prefix, template, func(obj model.Object) (stop bool) {
		count++
		assert.Equal(t, obj, template)
		return false
	})

	assert.NoError(t, err)
	assert.Equal(t, 1, count) // assert that the callback was called once
}

type testObject struct {
	Message string `validate:"required"`
}

func (to *testObject) MarshalBinary() (data []byte, err error) {
	return json.Marshal(to)
}

func (to *testObject) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, to)
}

func (to *testObject) Clone() model.Object {
	return &testObject{Message: to.Message}
}

// Validate should check the object's internal state for consistency or
// other correctness criteria. It should return an error if the validation fails.
func (to *testObject) Validate() error {
	return model.NewValidator().Struct(to)
}

var testObject1 = &testObject{Message: "Hello World"}

package inmem

import (
	"context"
	"testing"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/stretchr/testify/assert"
)

func TestKeyValueDB_Save(t *testing.T) {
	kv := new(KeyValueDB)
	ctx := context.Background()
	key := keyvalue.Key("testKey")
	value := keyvalue.Value("testValue")

	err := kv.Set(ctx, key, value)
	assert.NoError(t, err)

	loadedValue, err := kv.Get(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, value, loadedValue)
}

func TestKeyValueDB_Load(t *testing.T) {
	kv := new(KeyValueDB)
	ctx := context.Background()
	key := keyvalue.Key("testKey")
	value := keyvalue.Value("testValue")

	_, err := kv.Get(ctx, key)
	assert.Equal(t, keyvalue.ErrNotFound, err)

	_ = kv.Set(ctx, key, value)
	loadedValue, err := kv.Get(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, value, loadedValue)
}

func TestKeyValueDB_Delete(t *testing.T) {
	kv := new(KeyValueDB)
	ctx := context.Background()
	key := keyvalue.Key("testKey")
	value := keyvalue.Value("testValue")

	err := kv.Del(ctx, key)
	assert.NoError(t, err)

	_ = kv.Set(ctx, key, value)
	err = kv.Del(ctx, key)
	assert.NoError(t, err)

	_, err = kv.Get(ctx, key)
	assert.Equal(t, keyvalue.ErrNotFound, err)
}

func TestKeyValueDB_Iterate(t *testing.T) {
	kv := new(KeyValueDB)
	ctx := context.Background()

	_ = kv.Set(ctx, keyvalue.Key("testKey1"), keyvalue.Value("testValue1"))
	_ = kv.Set(ctx, keyvalue.Key("testKey2"), keyvalue.Value("testValue2"))
	_ = kv.Set(ctx, keyvalue.Key("otherKey"), keyvalue.Value("otherValue"))

	it, err := kv.Iter(ctx, keyvalue.Prefix("testKey"))
	assert.NoError(t, err)

	var count int
	for it.HasNext() {
		key, value, err := it.Next()
		assert.NoError(t, err)
		assert.True(
			t,
			key == "testKey1" && string(value) == "testValue1" ||
				key == "testKey2" && string(value) == "testValue2",
		)
		count++
	}
	assert.Equal(t, 2, count)

	err = it.Close()
	assert.NoError(t, err)
}

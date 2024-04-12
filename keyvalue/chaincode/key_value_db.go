package chaincode

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/anoideaopen/token/keyvalue"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

var _ keyvalue.DB = &KeyValueDB{}

// ErrChaincodeNilStub returns when ChaincodeStorage stub is empty.
var ErrChaincodeNilStub = errors.New("chaincode stub is nil")

// KeyValueDB is a structure that provides an interface to the storage of a chaincode.
// It includes a stub that provides various functions for interacting with that.
type KeyValueDB struct {
	Stub shim.ChaincodeStubInterface
}

// Set takes a context, a key, and a value, and saves the value associated with the key in the
// chaincode.
func (db *KeyValueDB) Set(_ context.Context, k keyvalue.Key, v keyvalue.Value) error {
	if db.Stub == nil {
		return internalError(ErrChaincodeNilStub)
	}

	compositeKey, err := db.tryComposite(k)
	if err != nil {
		return internalError(err)
	}

	if err := db.Stub.PutState(compositeKey, v); err != nil {
		return internalError(err)
	}

	return nil
}

// Get takes a context and a key, and returns the value associated with the key from the
// chaincode.
func (db *KeyValueDB) Get(_ context.Context, k keyvalue.Key) (keyvalue.Value, error) {
	if db.Stub == nil {
		return nil, internalError(ErrChaincodeNilStub)
	}

	compositeKey, err := db.tryComposite(k)
	if err != nil {
		return nil, internalError(err)
	}

	value, err := db.Stub.GetState(compositeKey)
	if err != nil {
		return nil, internalError(err)
	}

	if value == nil {
		return nil, keyvalue.ErrNotFound
	}

	return value, nil
}

// Del takes a context and a key, and deletes the value associated with the key from the
// chaincode.
func (db *KeyValueDB) Del(_ context.Context, k keyvalue.Key) error {
	if db.Stub == nil {
		return internalError(ErrChaincodeNilStub)
	}

	compositeKey, err := db.tryComposite(k)
	if err != nil {
		return internalError(err)
	}

	if err := db.Stub.DelState(compositeKey); err != nil {
		return internalError(err)
	}

	return nil
}

// Iter takes a context and a prefix, and returns an iterator over the keys in the chaincode
// storage that match the prefix.
func (db *KeyValueDB) Iter(_ context.Context, p keyvalue.Prefix) (keyvalue.Iterator, error) {
	if db.Stub == nil {
		return nil, internalError(ErrChaincodeNilStub)
	}

	var (
		iter shim.StateQueryIteratorInterface
		err  error
	)

	if keys := strings.Split(string(p), keyvalue.KeySeparator); len(keys) > 1 {
		iter, err = db.Stub.GetStateByPartialCompositeKey(keys[0], keys[1:])
	} else {
		iter, err = db.Stub.GetStateByPartialCompositeKey(string(p), []string{})
	}

	if err != nil {
		return nil, internalError(err)
	}

	return &chaincodeIterator{iter: iter}, nil
}

func (db *KeyValueDB) tryComposite(k keyvalue.Key) (string, error) {
	if keys := strings.Split(string(k), keyvalue.KeySeparator); len(keys) > 1 {
		key, err := db.Stub.CreateCompositeKey(keys[0], keys[1:])
		if err != nil {
			return "", err
		}

		return key, nil
	}

	return string(k), nil
}

// chaincodeIterator is a structure that provides an interface for iterating over the keys
// in the chaincode.
type chaincodeIterator struct {
	iter shim.StateQueryIteratorInterface
}

// HasNext checks whether the iterator has a next value.
func (i *chaincodeIterator) HasNext() bool {
	return i.iter.HasNext()
}

// Next returns the next key-value pair from the iterator.
func (i *chaincodeIterator) Next() (keyvalue.Key, keyvalue.Value, error) {
	response, err := i.iter.Next()
	if err != nil {
		return "", nil, internalError(err)
	}

	// https://github.com/hyperledger/fabric-chaincode-go/blob/main/shim/stub.go#L469
	const minUnicodeRuneValue = 0

	var (
		componentIndex = 1
		components     = []string{}
	)

	for i := 1; i < len(response.Key); i++ {
		if response.Key[i] == minUnicodeRuneValue {
			components = append(components, response.Key[componentIndex:i])
			componentIndex = i + 1
		}
	}

	return keyvalue.Key(strings.Join(components, keyvalue.KeySeparator)), response.Value, nil
}

// Close method that will release resources associated with the Iterator.
func (i *chaincodeIterator) Close() error {
	if err := i.iter.Close(); err != nil {
		return internalError(err)
	}
	return nil
}

func internalError(cause error) error {
	return fmt.Errorf("%w: %s", keyvalue.ErrInternal, cause.Error())
}

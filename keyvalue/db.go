package keyvalue

import (
	"context"
	"errors"
	"io"
	"strings"
)

// Predefined errors that may be returned by the functions in the storage package.
var (
	ErrNotFound = errors.New("key not found")
	ErrInternal = errors.New("internal error")
)

// KeySeparator symbol is used when you want to make a composite key. It stands for the path
// to the value, as if it were a directory file. This symbol is convenient to use for
// enumerating and searching for keys.
const KeySeparator = "/"

// Join creates a string, contains all substrings divided by KeySeparator.
// Example: "a" = "a"; "a", "b" = "a/b"; "a", "b", "c" = "a/b/c".
// The function ignores empty strings and KeySeparator.
func Join(s ...string) (out string) {
	var sb strings.Builder
	for _, ss := range s {
		if ss == "" || ss == KeySeparator {
			continue
		}
		sb.WriteString(KeySeparator)
		sb.WriteString(ss)
	}

	if out = sb.String(); out != "" {
		return out[1:]
	}

	return
}

type (
	// Key represents a unique identifier used to store and retrieve Value instances.
	Key string

	// Value contains the data associated with a Key in the storage, represented as bytes.
	Value []byte

	// Prefix is used to filter keys in the storage during iteration.
	Prefix string
)

// DB interface defines the basic methods for storing, retrieving,
// and iterating over data in a generic storage. This interface could be
// implemented in various contexts, including for testing purposes and in production environments.
//
//go:generate mockgen -source db.go -destination mock/mock_db.go -package mock
type DB interface {
	// Set method stores the provided Value under the given Key in the storage.
	// The Key uniquely identifies the Value for later retrieval.
	Set(context.Context, Key, Value) error

	// Get method retrieves the data associated with the provided Key from the storage.
	// If the Key does not exist in the storage, it returns an ErrKeyNotFound error.
	Get(context.Context, Key) (Value, error)

	// Del method removes the data associated with the provided Key from the storage.
	// If the Key does not exist in the storage anyway it returns nil.
	Del(context.Context, Key) error

	// Iter method returns an Iterator for keys with the provided Prefix in the storage.
	Iter(context.Context, Prefix) (Iterator, error)
}

// Iterator interface provides methods for iterating over keys and values in the storage.
type Iterator interface {
	// HasNext method returns a boolean indicating if there are more keys to iterate over.
	HasNext() bool

	// Next method returns the next key-value pair from the storage.
	// It returns an error if there are no more keys or if an issue occurs during retrieval.
	Next() (Key, Value, error)

	// The Iterator interface also includes the io.Closer interface,
	// requiring a Close method that will release resources associated with the Iterator.
	io.Closer
}

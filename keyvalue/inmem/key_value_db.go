package inmem

import (
	"context"
	"strings"
	"sync"

	"github.com/anoideaopen/token/keyvalue"
)

var _ keyvalue.DB = &KeyValueDB{}

// KeyValueDB implements a KeyValue interface using a map stored in memory.
type KeyValueDB struct {
	data map[keyvalue.Key]keyvalue.Value
	m    sync.RWMutex
}

// Set method stores the provided Value under the given Key in the
// The Key uniquely identifies the Value for later retrieval.
func (db *KeyValueDB) Set(_ context.Context, k keyvalue.Key, v keyvalue.Value) error {
	db.m.Lock()
	db.lazyInit()
	db.data[k] = v
	db.m.Unlock()

	return nil
}

// Get method retrieves the data associated with the provided Key from the
// If the Key does not exist in the storage, it returns an ErrKeyNotFound error.
func (db *KeyValueDB) Get(_ context.Context, k keyvalue.Key) (keyvalue.Value, error) {
	db.m.RLock()
	defer db.m.RUnlock()
	db.lazyInit()

	if v, ok := db.data[k]; ok {
		return v, nil
	}

	return nil, keyvalue.ErrNotFound
}

// Del method removes the data associated with the provided Key from the
// If the Key does not exist in the storage, it returns an ErrKeyNotFound error.
func (db *KeyValueDB) Del(_ context.Context, k keyvalue.Key) error {
	db.m.RLock()
	db.lazyInit()
	delete(db.data, k)
	db.m.RUnlock()

	return nil
}

// Iter method returns an Iterator for keys with the provided Prefix in the
func (db *KeyValueDB) Iter(_ context.Context, p keyvalue.Prefix) (keyvalue.Iterator, error) {
	db.m.RLock()
	defer db.m.RUnlock()
	db.lazyInit()

	i := new(inmemIter)
	for k, v := range db.data {
		if strings.HasPrefix(string(k), string(p)) {
			i.items = append(i.items, inmemItem{k: k, v: v})
		}
	}

	return i, nil
}

func (db *KeyValueDB) lazyInit() {
	if db.data == nil {
		db.data = make(map[keyvalue.Key]keyvalue.Value)
	}
}

// inmemIter interface provides methods for iterating over keys and values in the structure.
type inmemIter struct {
	items []inmemItem
	cur   int
	m     sync.Mutex
}

// inmemItem contains iteration inmemItem of slice.
type inmemItem struct {
	k keyvalue.Key
	v keyvalue.Value
}

// HasNext method returns a boolean indicating if there are more keys to iterate over.
func (i *inmemIter) HasNext() bool {
	i.m.Lock()
	defer i.m.Unlock()

	return i.cur != len(i.items)
}

// Next method returns the next key-value pair from the
// It returns an error if there are no more keys or if an issue occurs during retrieval.
func (i *inmemIter) Next() (keyvalue.Key, keyvalue.Value, error) {
	i.m.Lock()
	defer i.m.Unlock()

	if i.cur == len(i.items) {
		return "", nil, keyvalue.ErrNotFound
	}

	item := i.items[i.cur]
	i.cur++

	return item.k, item.v, nil
}

// The Iterator interface also includes the io.Closer interface,
// requiring a Close method that will release resources associated with the Iterator.
func (i *inmemIter) Close() error {
	i.m.Lock()
	defer i.m.Unlock()

	i.cur = 0
	i.items = nil

	return nil
}

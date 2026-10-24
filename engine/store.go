// Package engine defines the contract every Hoard storage backend implements.
//
// Every module in the curriculum builds toward a Store. The LSM tree, the
// B+tree, and the Raft-replicated store all satisfy this same interface, so
// the same tests and benchmarks run against all of them.
package engine

import "errors"

// ErrNotFound is returned by Get when a key does not exist or was deleted.
var ErrNotFound = errors.New("hoard: key not found")

// ErrClosed is returned by any operation on a Store after Close.
var ErrClosed = errors.New("hoard: store is closed")

// Store is an ordered key-value store. Keys are compared bytewise.
//
// Durability contract: once Put or Delete returns nil, the change must
// survive a process crash. Modules 1 and 2 are about earning that sentence.
type Store interface {
	// Get returns the value for key, or ErrNotFound.
	// The returned slice must not be modified by the caller.
	Get(key []byte) ([]byte, error)

	// Put sets key to value, overwriting any existing value.
	// The store must not retain references to key or value after returning.
	Put(key, value []byte) error

	// Delete removes key. Deleting a missing key is not an error.
	Delete(key []byte) error

	// Scan returns an iterator over keys in [start, end) in ascending order.
	// A nil start means "from the first key"; a nil end means "to the last key".
	Scan(start, end []byte) (Iterator, error)

	// Close flushes anything buffered and releases resources.
	Close() error
}

// Iterator walks a range of keys in ascending order.
//
// Usage:
//
//	it, err := s.Scan(nil, nil)
//	if err != nil { ... }
//	defer it.Close()
//	for it.Next() {
//		use(it.Key(), it.Value())
//	}
//	if err := it.Err(); err != nil { ... }
type Iterator interface {
	// Next advances to the next key. It returns false when the range is
	// exhausted or an error occurred; check Err to tell them apart.
	Next() bool
	// Key returns the current key. Valid only after Next returns true.
	Key() []byte
	// Value returns the current value. Valid only after Next returns true.
	Value() []byte
	// Err returns the first error encountered during iteration, if any.
	Err() error
	// Close releases the iterator.
	Close() error
}

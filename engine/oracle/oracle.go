// Package oracle is a deliberately simple, obviously correct in-memory Store.
//
// It exists to be the source of truth in tests: run the same random sequence
// of operations against the oracle and against your engine, and any
// difference is a bug in your engine. It is not durable and not fast, and it
// should never become either.
package oracle

import (
	"bytes"
	"sort"
	"sync"

	"github.com/tjb/hoard/engine"
)

// Store is an in-memory, map-backed engine.Store.
type Store struct {
	mu     sync.RWMutex
	data   map[string][]byte
	closed bool
}

var _ engine.Store = (*Store)(nil)

// New returns an empty oracle store.
func New() *Store {
	return &Store{data: make(map[string][]byte)}
}

func (s *Store) Get(key []byte) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.closed {
		return nil, engine.ErrClosed
	}
	v, ok := s.data[string(key)]
	if !ok {
		return nil, engine.ErrNotFound
	}
	return v, nil
}

func (s *Store) Put(key, value []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return engine.ErrClosed
	}
	s.data[string(key)] = bytes.Clone(value)
	return nil
}

func (s *Store) Delete(key []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed {
		return engine.ErrClosed
	}
	delete(s.data, string(key))
	return nil
}

func (s *Store) Scan(start, end []byte) (engine.Iterator, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.closed {
		return nil, engine.ErrClosed
	}
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		kb := []byte(k)
		if start != nil && bytes.Compare(kb, start) < 0 {
			continue
		}
		if end != nil && bytes.Compare(kb, end) >= 0 {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	vals := make([][]byte, len(keys))
	for i, k := range keys {
		vals[i] = s.data[k]
	}
	return &iterator{keys: keys, vals: vals, pos: -1}, nil
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closed = true
	return nil
}

// iterator holds a snapshot of the range taken when Scan was called.
type iterator struct {
	keys []string
	vals [][]byte
	pos  int
}

func (it *iterator) Next() bool {
	if it.pos+1 >= len(it.keys) {
		it.pos = len(it.keys)
		return false
	}
	it.pos++
	return true
}

func (it *iterator) Key() []byte   { return []byte(it.keys[it.pos]) }
func (it *iterator) Value() []byte { return it.vals[it.pos] }
func (it *iterator) Err() error    { return nil }
func (it *iterator) Close() error  { return nil }

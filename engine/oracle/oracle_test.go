package oracle

import (
	"errors"
	"testing"

	"github.com/tjb/hoard/engine"
)

func TestOracleBasics(t *testing.T) {
	s := New()

	if _, err := s.Get([]byte("missing")); !errors.Is(err, engine.ErrNotFound) {
		t.Fatalf("Get missing: got %v, want ErrNotFound", err)
	}

	for _, kv := range [][2]string{{"b", "2"}, {"a", "1"}, {"c", "3"}} {
		if err := s.Put([]byte(kv[0]), []byte(kv[1])); err != nil {
			t.Fatal(err)
		}
	}
	if err := s.Delete([]byte("c")); err != nil {
		t.Fatal(err)
	}

	it, err := s.Scan(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer it.Close()

	var got []string
	for it.Next() {
		got = append(got, string(it.Key())+"="+string(it.Value()))
	}
	want := []string{"a=1", "b=2"}
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Fatalf("Scan: got %v, want %v", got, want)
	}
}

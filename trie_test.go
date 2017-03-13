package trie

import (
	"strings"
	"testing"

	"github.com/satori/go.uuid"
)

func TestInsertWithHexadecimalCharSet(t *testing.T) {
	tree := NewTrie(HexadecimalCharSet)

	keys := []string{"abcd12365", "1234567890", "dfe507a000", "abcdef", "554654a654654b654654",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

	testInsert(tree, t, keys)
}

func TestInsertWithDecimalCharSet(t *testing.T) {
	tree := NewTrie(DecimalCharSet)
	keys := []string{"912365", "1234567890", "507000", "46546546546546", "554654654654654654",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "8", "7", "6", "5", "4", "3", "2"}

	testInsert(tree, t, keys)
}

func testInsert(tree *Trie, t *testing.T, keys []string) {
	set := map[string]bool{}

	for _, k := range keys {
		if !tree.Add(k) {
			t.Errorf("unexpected error inserting key: %s", k)
		}
		set[k] = true
	}

	for _, k := range keys {
		if !tree.Find(k) {
			t.Errorf("key %s not found", k)
		} else if !set[k] {
			t.Errorf("key %s not found in map", k)
		}
	}

	for _, bk := range []string{"a1b2b3c4d5f5e", "123", "ffffff", "aaa", "bcbe0123"} {
		if tree.Find(bk) {
			t.Errorf("key %s shouldn't be found", bk)
		} else if set[bk] {
			t.Errorf("key %s shouldn't found in map", bk)
		}
	}

	for _, k := range keys {
		if !tree.Delete(k) {
			t.Errorf("unexpected error deleting key: %s", k)
		}

		if tree.Find(k) {
			t.Errorf("key %s shouldn't found after deleting", k)
		}
	}
}

func BenchmarkUUIDMap(b *testing.B) {
	set := map[string]bool{}

	for n := 0; n < b.N; n++ {
		set[strings.Replace(uuid.NewV4().String(), "-", "", -1)] = true
	}
}

func BenchmarkUUIDTrie(b *testing.B) {
	tree := NewTrie(HexadecimalCharSet)

	for n := 0; n < b.N; n++ {
		tree.Add(strings.Replace(uuid.NewV4().String(), "-", "", -1))
	}
}

package trie

import "testing"

func TestInsert(t *testing.T) {
	tree := NewTrie(HexadecimalCharSet)

	keys := []string{"abcd12365", "1234567890", "dfe507a000", "abcdef", "554654a654654b654654",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}

	for _, k := range keys {
		if !tree.Add(k) {
			t.Errorf("unexpected error inserting key: %s", k)
		}
	}

	for _, k := range keys {
		if !tree.Find(k) {
			t.Errorf("key %s not found", k)
		}
	}

	for _, bk := range []string{"a1b2b3c4d5f5e", "123", "ffffff", "aaa", "bcbe0123"} {
		if tree.Find(bk) {
			t.Errorf("key %s shouldn't found", bk)
		}
	}
}

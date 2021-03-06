package trie

type Trie struct {
	endKey bool
	tree   []*Trie
	cs     CharSet
}

type CharSet interface {
	Size() int
	Position(byte) int
}

type hexCharSet struct{}

func (h *hexCharSet) Size() int {
	return 16
}

func (h *hexCharSet) Position(r byte) int {
	if r >= '0' && r <= '9' { // is a number
		return int(r - '0')
	}
	if r >= 'a' && r <= 'f' { // is a letter (a-f)
		return int(r-'a') + 10
	}

	return -1
}

type decimalCharSet struct{}

func (*decimalCharSet) Size() int {
	return 10
}

func (*decimalCharSet) Position(r byte) int {
	if r >= '0' && r <= '9' { // is a number
		return int(r - '0')
	}

	return -1
}

type alfaCharSet struct{}

func (*alfaCharSet) Size() int {
	return 26
}

func (*alfaCharSet) Position(r byte) int {
	if r >= 'a' && r <= 'z' { // is lower case
		return int(r - 'a')
	}
	if r >= 'A' && r <= 'Z' { // is upper case
		return int(r - 'A')
	}

	return -1
}

var (
	HexadecimalCharSet = &hexCharSet{}
	DecimalCharSet     = &decimalCharSet{}
	EnglishAlfaCharSet = &alfaCharSet{}
)

func NewTrie(cs CharSet) *Trie {
	return &Trie{
		tree: make([]*Trie, cs.Size()),
		cs:   cs,
	}
}

func (t *Trie) Add(str string) bool {
	l := len(str)
	if l == 0 {
		return false
	}

	for ; l > 0; l = len(str) {
		pos := t.cs.Position(str[0])
		if pos < 0 || pos > t.cs.Size() {
			// panic(pos)
			return false
		}

		if t.tree[pos] == nil {
			t.tree[pos] = NewTrie(t.cs)
		}
		if len(str) == 1 {
			t.tree[pos].endKey = true
			break
		}

		str = str[1:]
		t = t.tree[pos]
	}

	return true
}

func (t *Trie) Find(str string) bool {
	l := len(str)
	if l == 0 {
		return false
	}

	for ; l > 0; l = len(str) {
		pos := t.cs.Position(str[0])
		if pos < 0 || pos > len(t.tree) {
			// panic("incorrect position")
			return false
		}

		if t.tree[pos] == nil {
			return false
		}
		if l == 1 {
			return t.tree[pos].endKey
		}

		str = str[1:]
		t = t.tree[pos]
	}

	return false
}
func (t *Trie) Delete(str string) bool {
	l := len(str)
	if l == 0 {
		return false
	}

	for ; l > 0; l = len(str) {
		pos := t.cs.Position(str[0])
		if pos < 0 || pos > len(t.tree) {
			// panic("incorrect position")
			return false
		}

		if t.tree[pos] == nil {
			return false
		}
		if l == 1 {
			t.tree[pos].endKey = false
			return true
		}

		str = str[1:]
		t = t.tree[pos]
	}

	return false
}

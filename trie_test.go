package double_array

import (
	"testing"
	"bytes"
)

func TestTrieInsert(t *testing.T) {
	strs := []string{
		"aiueo",
		"あいうえお",
		"( ◠‿◠ )",
	}
	for _, str := range strs {
		trie := Trie{}
		trie.Insert(str)
		tree := trie.root
		buffer := bytes.NewBufferString(str)
		for {
			if char, err := buffer.ReadByte(); err != nil {
				break
			} else {
				if tree.children[0].key != char {
					t.Errorf("error")
				}
				tree = *tree.children[0]
			}
		}
	}
	trie := Trie{}
	trie.Insert(strs[0])
	trie.Insert(strs[1])
	if char, _ := bytes.NewBufferString("あ").ReadByte(); char != trie.root.children[1].key {
		t.Errorf("error")
	}
}

func TestTrieFind(t *testing.T) {
	trie := Trie{}
	trie.Insert("aiueo")

	errors := []string{
		"a",
		"aiueok",
		"x",
	}

	for _, val := range errors {
		if trie.Find(val) {
			t.Errorf("error")
		}
	}
	if !trie.Find("aiueo") {
		t.Errorf("error")
	}

}

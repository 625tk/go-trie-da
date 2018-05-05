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
		trie.insert(str)
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
	trie.insert(strs[0])
	trie.insert(strs[1])
	if char, _ := bytes.NewBufferString("あ").ReadByte(); char != trie.root.children[1].key {
		t.Errorf("error")
	}
}

func TestTrieFind(t *testing.T) {
	trie := Trie{}
	trie.insert("aiueo")

	errors := []string{
		"a",
		"aiueok",
		"x",
	}

	for _, val := range errors {
		if trie.find(val) {
			t.Errorf("error")
		}
	}
	if !trie.find("aiueo") {
		t.Errorf("error")
	}

}

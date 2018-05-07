package double_array

import (
	"testing"
	"bytes"
)

func TestTrieNode_Len(t *testing.T) {
	trie := TrieNode{}
	if trie.children.Len() != 0{
		t.Errorf("len error %d", trie.children.Len())
	}
	trie.children  = append(trie.children, &(TrieNode{}))
	if trie.children.Len() != 1{
		t.Errorf("len error %d", trie.children.Len())
	}
}

func TestTrieNode_Swap(t *testing.T) {
	trie := TrieNode{}
	trie.children  = append(trie.children, &(TrieNode{id: 0}))
	trie.children  = append(trie.children, &(TrieNode{id: 1}))
	trie.children.Swap(0,1)
	if trie.children[0].id != 1{
		t.Errorf("swap error %d", trie.children[0].id)
	}
	if trie.children[1].id != 0{
		t.Errorf("swap error %d", trie.children[1].id)
	}
}

func TestTrieNode_Less(t *testing.T) {
	trie := TrieNode{}
	trie.children  = append(trie.children, &(TrieNode{id: 0, key: byte(90)}))
	trie.children  = append(trie.children, &(TrieNode{id: 1, key: byte(80)}))
	trie.children  = append(trie.children, &(TrieNode{id: 1, key: byte(100)}))
	if trie.children.Less(0, 1){
		t.Errorf("less error %d, %d", 0, 1)
	}
	if !trie.children.Less(0, 2){
		t.Errorf("less error %d, %d", 0, 2)
	}
}

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

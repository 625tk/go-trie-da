package double_array

import (
	"bytes"
)

type (
	Trie struct {
		root TrieNode
		size int
	}

	TrieNode struct {
		children []*TrieNode
		id       int
		key      byte
		value    int
	}
)

func (trie *Trie) insert(str string) {
	tree := &trie.root

	buffer := bytes.NewBufferString(str)
	for {
		if elm, err := buffer.ReadByte(); err != nil {
			break
		} else {
			index := -1
			for i, child := range tree.children {
				if child.key == elm {
					index = i
				}
			}
			if index == -1 {
				index = len(tree.children)
				tree.children = append(tree.children, &TrieNode{id: trie.size, key: elm})
				trie.size++
			}
			tree = tree.children[index]
		}
	}
	tree.value = 1
}

func (trie *Trie) find(str string) bool {
	buffer := bytes.NewBufferString(str)
	tree := &trie.root
	for {
		if elm, err := buffer.ReadByte(); err != nil {
			break
		} else {
			index := -1
			for i, child := range tree.children {
				if child.key == elm {
					index = i
				}
			}
			if index == -1 {
				return false
			}
			tree = tree.children[index]
		}
	}
	return tree.value > 0
}

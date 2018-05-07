package double_array

import (
	"bytes"
)

type (
	Trie struct {
		root TrieNode
		Size int
	}

	TrieNodes []*TrieNode

	TrieNode struct {
		children TrieNodes
		id       int
		key      byte
		value    int
	}
)

func (t TrieNodes) Len() int {
	return len(t)
}

func (t TrieNodes) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TrieNodes) Less(i, j int) bool {
	return t[i].key < t[j].key
}

func (trie *Trie) Insert(str string) {
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
				tree.children = append(tree.children, &TrieNode{id: trie.Size, key: elm})
				trie.Size++
			}
			tree = tree.children[index]
		}
	}
	tree.value = tree.value + 1
}

func (trie *Trie) Find(str string) bool {
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

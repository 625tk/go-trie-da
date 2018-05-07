package double_array

import (
	"bytes"
	"fmt"
)

type DoubleArray []struct {
	BASE  int
	CHECK int
	VALUE int
}

var SIZE_MAX int = 1000000

func (DoubleArray) Build(trie *Trie) *DoubleArray {
	da := make(DoubleArray, SIZE_MAX)
	// init
	for i := 0; i < SIZE_MAX; i++ {
		da[i].BASE = -1
		da[i].CHECK = -1
	}
	da[0].BASE = 0
	head := 1
	da[0].BASE = da.insert(&head, &trie.root, 0)
	return &da
}

func (da *DoubleArray) insert(head *int, tree *TrieNode, check int) int {
	if tree.children.Len() == 0 {
		return -1
	}
	min_key := byte(255)
	for _, child := range tree.children {
		if min_key > child.key {
			min_key = child.key
		}
	}

	ok := false
	empty_slot := *head

	for !ok {
		ok = true
		offset := empty_slot - int(min_key)
		if offset < 0 {
			offset += int(min_key)
		}

		for _, child := range tree.children {
			ok = ok && ((*da)[int(child.key)+offset].CHECK < 0)
			if !ok {
				break;
			}
		}

		if ok {
			for _, child := range tree.children {
				current_slot := int(child.key) + offset
				prev_empty_slot := current_slot + (*da)[current_slot].BASE
				next_empty_slot := current_slot - (*da)[current_slot].CHECK

				if (*da)[current_slot].BASE != 0 {
					(*da)[prev_empty_slot].CHECK += (*da)[current_slot].CHECK
					(*da)[next_empty_slot].BASE += (*da)[current_slot].BASE
				} else {
					// update head
					*head = next_empty_slot
					(*da)[next_empty_slot].BASE = 0
				}

				(*da)[current_slot].CHECK = check
				(*da)[current_slot].VALUE = child.value
			}

			for _, child := range (tree.children) {
				(*da)[int(child.key)+offset].BASE = da.insert(head, child, int(child.key)+offset)
			}
			return offset
		}
		//update offset
		empty_slot -= (*da)[empty_slot].CHECK
	}
	return -1
}

func (da *DoubleArray) Query(str string) int {
	head := 0

	buffer := bytes.NewBufferString(str)
	for {
		char, err := buffer.ReadByte()
		if err != nil {
			if val := (*da)[head].VALUE; val == 0 {
				return -1
			} else {
				return val
			}
		}
		next := int(char) + (*da)[head].BASE
		if next < -1 || (*da)[next].CHECK != head {
			return -1
		}
		head = next
	}
	return -1
}

func (da *DoubleArray) Show() {
	for i := 0; i < 140; i++ {
		fmt.Printf("%3d", i)
	}
	fmt.Println()
	for i := 0; i < 140; i++ {
		fmt.Printf("%3d", (*da)[i].BASE)
	}
	fmt.Println()
	for i := 0; i < 140; i++ {
		fmt.Printf("%3d", (*da)[i].CHECK)
	}
	fmt.Println()
	for i := 0; i < 140; i++ {
		fmt.Printf("%3d", (*da)[i].VALUE)
	}
	fmt.Println()
}

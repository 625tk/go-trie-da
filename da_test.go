package double_array

import (
	"testing"
)

func TestDoubleArray_Build(t *testing.T) {
	trie := Trie{}
	trie.Insert("A")
	trie.Insert("AA")
	trie.Insert("B")
	da := DoubleArray{}.Build(&trie)
	if (*da)['A'+1].CHECK != 0{
		t.Errorf("value error")
	}
	if (*da)['B'+1].CHECK != 0{
		t.Errorf("value error")
	}
	if (*da)['A'+3].CHECK != 'A'+1{
		t.Errorf("value error")
	}
}
func TestDoubleArray_Build2(t *testing.T) {
	trie := Trie{}
	trie.Insert("A")
	trie.Insert("A")
	trie.Insert("B")
	da := DoubleArray{}.Build(&trie)
	if (*da)['A'+1].CHECK != 0{
		t.Errorf("value error")
	}
	if (*da)['B'+1].CHECK != 0{
		t.Errorf("value error")
	}
}

func TestDoubleArray_Query(t *testing.T) {
	trie := Trie{}
	trie.Insert("aiueo")
	trie.Insert("aiueok")
	da := DoubleArray{}.Build(&trie)

	if da.Query("aiueo") != 1 {
		t.Errorf("query failed")
	}
	if da.Query("aiueok") != 1 {
		t.Errorf("query failed")
	}

	if da.Query("aiue") != -1 {
		t.Errorf("query failed")
	}

	if da.Query("aiueoka") != -1 {
		t.Errorf("query failed")
	}
}

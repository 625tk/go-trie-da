
# go double array

## create trie

```go:
trie := Trie{}
trie.Insert("abc")
trie.Insert("def")
trie.Insert("abcd")

trie.Find("abcd")
```


## create double array from trie

```go:
trie := Trie{}
trie.Insert("abc")
trie.Insert("def")
trie.Insert("abcd")

da := DoubleArray{}.Build(&trie)
da.Query("abcd")
da.Query("def")

```

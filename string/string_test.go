package string

import (
	"fmt"
	"testing"
)

var trie *Trie

func init() {
	trie = NewTrie()
	trie.Add("abc")
	trie.Add("bce")
	trie.Add("abe")
}

func Test_TrieSearch(t *testing.T) {
	fmt.Println(trie.Search("ab"))
	fmt.Println(trie.StartWith("ba"))
}

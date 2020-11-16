package string

type Node struct {
	Next   map[rune]*Node
	IsWord bool
}

type Trie struct {
	root *Node
}

func newNode() *Node {
	return &Node{
		Next: make(map[rune]*Node),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) Add(word string) {
	lastNode := t.root
	for _, r := range []rune(word) {
		if n, ok := lastNode.Next[r]; ok {
			lastNode = n
		} else {
			node := newNode()
			lastNode.Next[r] = node
			lastNode = node
		}
	}
	lastNode.IsWord = true
}

func (t *Trie) Search(word string) bool {
	lastNode := t.root
	for _, r := range []rune(word) {
		if n, ok := lastNode.Next[r]; ok {
			lastNode = n
		} else {
			return false
		}
	}
	return lastNode.IsWord
}

func (t *Trie) StartWith(word string) bool {
	lastNode := t.root
	for _, r := range []rune(word) {
		if n, ok := lastNode.Next[r]; ok {
			lastNode = n
		} else {
			return false
		}
	}
	return true
}

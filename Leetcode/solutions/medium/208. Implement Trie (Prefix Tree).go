package medium

import "fmt"

// Reference: https://leetcode.com/problems/implement-trie-prefix-tree/
func init() {
	Solutions[208] = func() {
		fmt.Println("Input: ['Trie', 'insert', 'search', 'search', 'startsWith', 'insert', 'search']")
		fmt.Println("       [[], ['apple'], ['apple'], ['app'], ['app'], ['app'], ['app']]")
		fmt.Println("Output:")

		trie := TrieConstructor()
		trie.Insert("apple")
		fmt.Println("trie.Search('apple')", trie.Search("apple"))
		fmt.Println("trie.Search('app')", trie.Search("app"))
		fmt.Println("trie.StartsWith('app')", trie.StartsWith("app"))
		trie.Insert("app")
		fmt.Println("trie.Search('app')", trie.Search("app"))
	}
}

type Trie struct {
	root *TrieNodes
}

func (t *Trie) init() {
	t.root = new(TrieNodes)
}

func TrieConstructor() Trie {
	trie := Trie{}
	trie.init()
	return trie
}

func (t *Trie) Insert(word string) {
	current := t.root
	for i := range word {
		if current.Get(word[i]) == nil {
			current.Set(word[i], new(TrieNodes))
		}
		current = current.Get(word[i])
	}
	current.Tail = true
}

func (t *Trie) Search(word string) bool {
	node := t.ToTail(word)
	if node != nil {
		return node.Tail
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.ToTail(prefix) != nil
}

func (t *Trie) ToTail(prefix string) *TrieNodes {
	current := t.root
	for i := range prefix {
		current = current.Get(prefix[i])
		if current == nil {
			return current
		}
	}
	return current
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/implement-trie-prefix-tree/
func Leetcode_Trie_Constructor() {
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

type TrieNodes struct {
	suffixes [26]*TrieNodes
	Tail     bool
}

func (tr *TrieNodes) get(ch byte) *TrieNodes {
	return tr.suffixes[ch-'a']
}

func (tr *TrieNodes) set(ch byte, value *TrieNodes) {
	tr.suffixes[ch-'a'] = value
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
		if current.get(word[i]) == nil {
			current.set(word[i], new(TrieNodes))
		}
		current = current.get(word[i])
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
		current = current.get(prefix[i])
		if current == nil {
			return current
		}
	}
	return current
}

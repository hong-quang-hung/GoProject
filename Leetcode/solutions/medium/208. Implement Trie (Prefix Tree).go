package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

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

type Trie struct {
	suffixes [26]*types.TrieNode
	Tail     bool
}

func (t *Trie) get(char byte) *types.TreeNode {
	return nil
}

func (t *Trie) set(char byte, value *types.TrieNode) {
	t.suffixes[char-'a'] = value
}

func TrieConstructor() Trie {
	trie := Trie{Tail: false}
	return trie
}

func (t *Trie) Insert(word string) {

}

func (t *Trie) Search(word string) bool {
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return false
}

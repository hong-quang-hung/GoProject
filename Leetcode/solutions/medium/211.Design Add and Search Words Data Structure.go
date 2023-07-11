package medium

import "fmt"

// Reference: https://leetcode.com/problems/design-add-and-search-words-data-structure/
func init() {
	Solutions[211] = func() {
		wordDictionary := Constructor()
		wordDictionary.AddWord("bad")
		wordDictionary.AddWord("dad")
		wordDictionary.AddWord("mad")
		fmt.Println(wordDictionary.Search("pad"))
		fmt.Println(wordDictionary.Search("bad"))
		fmt.Println(wordDictionary.Search(".ad"))
		fmt.Println(wordDictionary.Search("b.."))
	}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	wd := WordDictionary{}
	wd.Init()
	return wd
}

func (wd *WordDictionary) Init() {
	wd.root = new(TrieNode)
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd.root
	for i := range word {
		ch := word[i]
		if node.Childrens == nil {
			node.Childrens = make(map[byte]*TrieNode)
		}
		if node.Childrens[ch] == nil {
			node.Childrens[ch] = new(TrieNode)
		}
		node = node.Childrens[ch]
	}
	node.End = true
}

func (wd *WordDictionary) Search(word string) bool {
	return helper(wd.root, word, 0)
}

func helper(root *TrieNode, word string, i int) bool {
	for i < len(word) {
		ch := word[i]
		if ch == '.' {
			for c := range root.Childrens {
				root := root.Childrens[c]
				if helper(root, word, i+1) {
					return true
				}
			}
			return false
		} else {
			if root.Childrens[ch] == nil {
				return false
			}
			root = root.Childrens[ch]
		}
		i++
	}
	return root.End
}

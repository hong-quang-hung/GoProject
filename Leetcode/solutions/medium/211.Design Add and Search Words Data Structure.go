package solutions

import "leetcode.com/Leetcode/types"

// Reference: https://leetcode.com/problems/symmetric-tree/
type WordDictionary struct {
	root *types.TrieNode
}

func Constructor() WordDictionary {
	wd := WordDictionary{}
	wd.Init()
	return wd
}

func (wd *WordDictionary) Init() {
	wd.root = new(types.TrieNode)
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd.root
	for i := range word {
		ch := word[i]
		if node.Childrens == nil {
			node.Childrens = make(map[byte]*types.TrieNode)
		}
		if node.Childrens[ch] == nil {
			node.Childrens[ch] = new(types.TrieNode)
		}
		node = node.Childrens[ch]
	}
	node.End = true
}

func (wd *WordDictionary) Search(word string) bool {
	return helper(wd.root, word, 0)
}

func helper(root *types.TrieNode, word string, i int) bool {
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

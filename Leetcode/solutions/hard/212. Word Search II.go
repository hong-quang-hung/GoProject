package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/word-search-ii/
func init() {
	Solutions[212] = func() {
		fmt.Println(`Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]`)
		fmt.Println(`Output:`, findWords(S2SoSliceByte(`[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]`), []string{"oath", "pea", "eat", "rain"}))
		fmt.Println(`Input: board = [["a","b"],["c","d"]], words = ["abcb"]`)
		fmt.Println(`Output:`, findWords(S2SoSliceByte(`[["a","b"],["c","d"]]`), []string{"abcb"}))
	}
}

type WordSearchTrieNode struct {
	data map[byte]*WordSearchTrieNode
	tail bool
}

type WordSearchTrie struct {
	root *WordSearchTrieNode
}

func NewWordSearchTrie() *WordSearchTrie {
	return &WordSearchTrie{root: &WordSearchTrieNode{data: make(map[byte]*WordSearchTrieNode)}}
}

func (t *WordSearchTrie) InsertWord(word string) {
	current := t.root
	for i := range word {
		if _, ok := current.data[word[i]]; !ok {
			current.data[word[i]] = &WordSearchTrieNode{data: make(map[byte]*WordSearchTrieNode)}
		}
		current = current.data[word[i]]
	}
	current.tail = true
}

func findWords(board [][]byte, words []string) []string {
	trie := NewWordSearchTrie()
	for _, word := range words {
		trie.InsertWord(word)
	}

	m, n := len(board), len(board[0])
	wordsFound := make(map[string]interface{}, 0)

	var f func(r, c int, s string, node *WordSearchTrieNode)
	f = func(r, c int, s string, node *WordSearchTrieNode) {
		if node.tail {
			wordsFound[s] = struct{}{}
			node.tail = false
		}

		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] == '*' {
			return
		}

		if _, ok := node.data[board[r][c]]; !ok {
			return
		}

		next := s + string(board[r][c])
		visited := board[r][c]
		board[r][c] = '*'
		f(r-1, c, next, node.data[visited])
		f(r+1, c, next, node.data[visited])
		f(r, c-1, next, node.data[visited])
		f(r, c+1, next, node.data[visited])
		board[r][c] = visited
	}

	res := make([]string, 0, len(wordsFound))
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f(i, j, "", trie.root)
		}
	}

	for word := range wordsFound {
		res = append(res, word)
	}
	return res
}

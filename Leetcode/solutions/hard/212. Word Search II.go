package hard

import "fmt"

// Reference: https://leetcode.com/problems/word-search-ii/
func init() {
	Solutions[212] = func() {
		fmt.Println(`Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]`)
		fmt.Println(`Output:`, findWords(S2SoSliceByte(`[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]`), []string{"oath", "pea", "eat", "rain"}))
		fmt.Println(`Input: board = [["a","b"],["c","d"]], words = ["abcb"]`)
		fmt.Println(`Output:`, findWords(S2SoSliceByte(`[["a","b"],["c","d"]]`), []string{"abcb"}))
	}
}

func findWords(board [][]byte, words []string) []string {
	return nil
}

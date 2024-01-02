package hard

import "fmt"

// Reference: https://leetcode.com/problems/word-ladder/
func init() {
	Solutions[127] = func() {
		fmt.Println(`Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]`)
		fmt.Println(`Output:`, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
		fmt.Println(`Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]`)
		fmt.Println(`Output:`, ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	}
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wl := make(map[string]interface{})
	for _, word := range wordList {
		wl[word] = struct{}{}
	}

	n := len(beginWord)
	visited := make(map[string]interface{})
	visited[beginWord] = struct{}{}
	queue := []string{beginWord}
	res := 1
	for len(queue) > 0 {
		nextWords := []string{}
		for _, q := range queue {
			if q == endWord {
				return res
			}

			for i := 0; i < n; i++ {
				for j := 0; j < 26; j++ {
					nextWord := q[:i] + string(rune(97+j)) + q[i+1:]
					if _, isVisited := visited[nextWord]; !isVisited {
						if _, ok := wl[nextWord]; ok {
							nextWords = append(nextWords, nextWord)
							visited[nextWord] = struct{}{}
						}
					}
				}
			}
		}

		queue = nextWords
		nextWords = nil
		res++
	}
	return 0
}

package hard

import "fmt"

// Reference: https://leetcode.com/problems/edit-distance/
func init() {
	Solutions[72] = func() {
		fmt.Println(`Input: word1 = "intention", word2 = "execution"`)
		fmt.Println(`Output:`, minDistance(`intention`, `execution`))
	}
}

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1)+1, len(word2)+1
	table := make([][]int, len1)
	table[0] = make([]int, len2)
	for w1 := 1; w1 < len1; w1++ {
		table[w1] = make([]int, len2)
		table[w1][0] = w1
	}

	for w2 := 1; w2 < len2; w2++ {
		table[0][w2] = w2
	}

	for w1 := 1; w1 < len1; w1++ {
		for w2 := 1; w2 < len2; w2++ {
			if word2[w2-1] == word1[w1-1] {
				table[w1][w2] = table[w1-1][w2-1]
			} else {
				table[w1][w2] = min(table[w1-1][w2], min(table[w1][w2-1], table[w1-1][w2-1])) + 1
			}
		}
	}
	return table[len1-1][len2-1]
}

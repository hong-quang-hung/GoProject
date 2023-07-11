package medium

import "fmt"

// Reference: https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/
func init() {
	Solutions[2038] = func() {
		fmt.Println("Input: colors = 'AAABABB'")
		fmt.Println("Output:", winnerOfGame("AAABABB"))
		fmt.Println("Input: colors = 'AA'")
		fmt.Println("Output:", winnerOfGame("AA"))
		fmt.Println("Input: colors = 'ABBBBBBBAAA'")
		fmt.Println("Output:", winnerOfGame("ABBBBBBBAAA"))
	}
}

func winnerOfGame(colors string) bool {
	A, B := 0, 0
	for i := 1; i < len(colors)-1; i++ {
		if colors[i-1] == colors[i] && colors[i+1] == colors[i] {
			if colors[i] == 'A' {
				A++
			} else {
				B++
			}
		}
	}
	return A > B
}

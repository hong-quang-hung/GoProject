package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/determine-the-winner-of-a-bowling-game/
func Leetcode_Is_Winner() {
	fmt.Println("Input: player1 = [4,10,7,9], player2 = [6,5,2,3]")
	fmt.Println("Output:", isWinner([]int{4, 10, 7, 9}, []int{6, 5, 2, 3}))
	fmt.Println("Input: player1 = [5,6,1,10], player2 = [5,1,10,5]")
	fmt.Println("Output:", isWinner([]int{5, 6, 1, 10}, []int{5, 1, 10, 5}))
	fmt.Println("Input: player1 = [10,2,2,3], player2 = [3,8,4,5]")
	fmt.Println("Output:", isWinner([]int{10, 2, 2, 3}, []int{3, 8, 4, 5}))
}

func isWinner(player1 []int, player2 []int) int {
	score1, score2, flag1, flag2 := 0, 0, -1, -1
	for i := 0; i < len(player1); i++ {
		if i <= flag1 {
			score1 += 2 * player1[i]
		} else {
			score1 += player1[i]
		}

		if i <= flag2 {
			score2 += 2 * player2[i]
		} else {
			score2 += player2[i]
		}

		if player1[i] == 10 {
			flag1 = i + 2
		}
		if player2[i] == 10 {
			flag2 = i + 2
		}
	}

	if score1 > score2 {
		return 1
	}

	if score1 < score2 {
		return 2
	}
	return 0
}

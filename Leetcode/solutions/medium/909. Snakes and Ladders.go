package medium

import "fmt"

// Reference: https://leetcode.com/problems/snakes-and-ladders/
func init() {
	Solutions[909] = func() {
		fmt.Println(`Input: board = [[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]`)
		fmt.Println(`Output:`, snakesAndLadders(S2SoSliceInt(`[[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]`)))
		fmt.Println(`Input: board = [[-1,-1],[-1,3]]`)
		fmt.Println(`Output:`, snakesAndLadders(S2SoSliceInt(`[[-1,-1],[-1,3]]`)))
	}
}

func snakesAndLadders(board [][]int) int {
	n := len(board)
	des := n * n
	dp := make([]int, des+1)
	dp[1] = 0
	for i := 2; i <= des; i++ {
		dp[i] = des
	}

	move := make([]int, des+1)
	count, cur := 0, 1
	for i := n - 1; i >= 0; i-- {
		if count == 0 {
			for count < n {
				move[cur] = board[i][count]
				cur++
				count++
			}
		} else {
			for count > 0 {
				count--
				move[cur] = board[i][count]
				cur++
			}
		}
	}

	queue := []int{1}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for j := i + 1; j <= min(i+6, n*n); j++ {
			if move[j] > 0 {
				if dp[move[j]] > dp[i]+1 {
					queue = append(queue, move[j])
					dp[move[j]] = dp[i] + 1
				}
				continue
			}
			if dp[j] > dp[i]+1 {
				queue = append(queue, j)
				dp[j] = dp[i] + 1
			}
		}
	}

	if dp[des] == des {
		return -1
	}
	return dp[des]
}

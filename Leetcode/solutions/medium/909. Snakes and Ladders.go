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
	return 0
}

package medium

import "fmt"

func init() {
	Solutions[2352] = Leetcode_Equal_Pairs
}

// Reference: https://leetcode.com/problems/equal-row-and-column-pairs/
func Leetcode_Equal_Pairs() {
	fmt.Println("Input: grid = [[3,2,1],[1,7,6],[2,7,7]]")
	fmt.Println("Output:", equalPairs(S2SoSliceInt("[[3,2,1],[1,7,6],[2,7,7]]")))
	fmt.Println("Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]")
	fmt.Println("Output:", equalPairs(S2SoSliceInt("[[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]")))
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	m := make(map[[200]int]int)
	for i := 0; i < n; i++ {
		key := [200]int{}
		for j := 0; j < n; j++ {
			key[j] = grid[i][j]
		}
		m[key]++
	}

	res := 0
	for i := 0; i < n; i++ {
		key := [200]int{}
		for j := 0; j < n; j++ {
			key[j] = grid[j][i]
		}
		res += m[key]
	}
	return res
}

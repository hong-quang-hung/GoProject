package easy

import "fmt"

func init() {
	Solutions[1351] = Leetcode_Count_Negatives
}

// Reference: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
func Leetcode_Count_Negatives() {
	fmt.Println("Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]")
	fmt.Println("Output:", countNegatives(S2SoSliceInt("[[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]")))
	fmt.Println("Input: grid = [[3,2],[1,0]]")
	fmt.Println("Output:", countNegatives(S2SoSliceInt("[[3,2],[1,0]]")))
}

func countNegatives(grid [][]int) int {
	n, m, res := len(grid), len(grid[0]), 0
	p := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= p; j-- {
			if grid[i][j] < 0 {
				res++
			} else {
				p = j
				break
			}
		}
	}
	return res
}

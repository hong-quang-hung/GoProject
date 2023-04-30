package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/
func Leetcode_Find_Max_Fish() {
	fmt.Println("Input: grid = [[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]")
	fmt.Println("Output:", findMaxFish(utils.S2SoSliceInt("[[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]")))
	fmt.Println("Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]")
	fmt.Println("Output:", findMaxFish(utils.S2SoSliceInt("[[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]")))
}

func findMaxFish(grid [][]int) int {
	return 0
}

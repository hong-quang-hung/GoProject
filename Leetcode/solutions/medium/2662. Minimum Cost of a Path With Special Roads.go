package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/minimum-cost-of-a-path-with-special-roads/
func Leetcode_Minimum_Cost() {
	fmt.Println("Input: start = [1,1], target = [4,5], specialRoads = [[1,2,3,3,2],[3,4,4,5,1]]")
	fmt.Println("Output:", minimumCost([]int{1, 1}, []int{4, 5}, utils.S2SoSliceInt("[[1,2,3,3,2],[3,4,4,5,1]]")))
	fmt.Println("Input: start = [3,2], target = [5,7], specialRoads = [[3,2,3,4,4],[3,3,5,5,5],[3,4,5,6,6]]")
	fmt.Println("Output:", minimumCost([]int{3, 2}, []int{5, 7}, utils.S2SoSliceInt("[[3,2,3,4,4],[3,3,5,5,5],[3,4,5,6,6]]")))
}

func minimumCost(start []int, target []int, specialRoads [][]int) int {
	dp := make([][]int, target[0]+1)
	for i := start[0]; i < target[0]+1; i++ {
		dp[i] = make([]int, target[1]+1)
		for j := start[1]; j < target[1]+1; j++ {
			dp[i][j] = -1
		}
	}

	m := make([][][]int, target[0]+1)
	for i := start[0]; i < target[0]+1; i++ {
		m[i] = make([][]int, target[1]+1)
	}

	for i := 0; i < len(specialRoads); i++ {
		m[specialRoads[i][2]][specialRoads[i][3]] = []int{specialRoads[i][0], specialRoads[i][1], specialRoads[i][4]}
	}
	return minimumCostSolved(start, target, m, dp, target[0], target[1])
}

func minimumCostSolved(start []int, target []int, m [][][]int, dp [][]int, i int, j int) int {
	if i < start[0] || j < start[1] {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	startInfo := m[i][j]
	if startInfo != nil && startInfo[0] != 0 && startInfo[1] != 0 {
		dp[i][j] = min(min(minimumCostSolved(start, target, m, dp, i-1, j)+1, minimumCostSolved(start, target, m, dp, i, j-1)+1), minimumCostSolved(start, target, m, dp, startInfo[0], startInfo[1])+startInfo[2])
	} else {
		dp[i][j] = min(minimumCostSolved(start, target, m, dp, i-1, j)+1, minimumCostSolved(start, target, m, dp, i, j-1)+1)
	}
	return dp[i][j]
}

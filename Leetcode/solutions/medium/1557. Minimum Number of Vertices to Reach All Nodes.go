package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
func Leetcode_Find_Smallest_Set_Of_Vertices() {
	fmt.Println("Input: n = 6, edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]")
	fmt.Println("Output:", findSmallestSetOfVertices(6, utils.S2SoSliceInt("[[0,1],[0,2],[2,5],[3,4],[4,2]]")))
	fmt.Println("Input: n = 5, edges = [[0,1],[2,1],[3,1],[1,4],[2,4]]")
	fmt.Println("Output:", findSmallestSetOfVertices(5, utils.S2SoSliceInt("[[0,1],[2,1],[3,1],[1,4],[2,4]]")))
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	res := []int{}
	exists := make(map[int]bool)
	for _, edge := range edges {
		exists[edge[1]] = true
	}

	for i := 0; i < n; i++ {
		if !exists[i] {
			res = append(res, i)
		}
	}
	return res
}
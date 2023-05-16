package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/combination-sum-ii/
func Leetcode_Combination_Sum_2() {
	fmt.Println("Input: candidates = [10,1,2,7,6,1,5], target = 8")
	fmt.Println("Output:", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println("Input: candidates = [2,5,2,1,2], target = 5")
	fmt.Println("Output:", combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	combinationSum2Backtrack(&res, candidates, target, []int{}, 0)
	return res
}

func combinationSum2Backtrack(res *[][]int, candidates []int, target int, current []int, i int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
		return
	}

	for j := i; j < len(candidates); j++ {
		if j > i && candidates[j] == candidates[j-1] {
			continue
		}

		if candidates[j] <= target {
			current = append(current, candidates[j])
			combinationSum2Backtrack(res, candidates, target-candidates[j], current, j+1)
			current = current[:len(current)-1]
		}
	}
}

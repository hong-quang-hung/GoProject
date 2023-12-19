package medium

import "fmt"

// Reference: https://leetcode.com/problems/combination-sum/
func init() {
	Solutions[39] = func() {
		fmt.Println(`Input: candidates = [2,3,6,7], target = 7`)
		fmt.Println(`Output:`, combinationSum([]int{2, 3, 6, 7}, 7))
		fmt.Println(`Input: candidates = [2,3,5], target = 8`)
		fmt.Println(`Output:`, combinationSum([]int{2, 3, 5}, 8))
		fmt.Println(`Input: candidates = [8,7,4,3], target = 11`)
		fmt.Println(`Output:`, combinationSum([]int{8, 7, 4, 3}, 11))
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	combinationSumDFS(&res, candidates, target, []int{}, 0)
	return res
}

func combinationSumDFS(res *[][]int, candidates []int, target int, current []int, i int) {
	if target == 0 {
		*res = append(*res, current)
	} else {
		for j := i; j < len(candidates); j++ {
			if candidates[j] <= target {
				temp := make([]int, len(current))
				copy(temp, current)
				current = append(current, candidates[j])
				combinationSumDFS(res, candidates, target-candidates[j], current, j)
				current = temp
			}
		}
	}
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/generate-parentheses/
func Leetcode_Combination_Sum() {
	fmt.Println("Input: candidates = [2,3,6,7], target = 7")
	fmt.Println("Output:", combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println("Input: candidates = [2,3,5], target = 8")
	fmt.Println("Output:", combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	return nil
}

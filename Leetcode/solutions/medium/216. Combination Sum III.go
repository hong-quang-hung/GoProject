package medium

import "fmt"

// Reference: https://leetcode.com/problems/combination-sum-iii/
func init() {
	Solutions[216] = func() {
		fmt.Println("Input: k = 3, n = 7")
		fmt.Println("Output:", combinationSum3(3, 7))
		fmt.Println("Input: k = 3, n = 9")
		fmt.Println("Output:", combinationSum3(3, 9))
		fmt.Println("Input: k = 2, n = 18")
		fmt.Println("Output:", combinationSum3(2, 18))
	}
}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	combinationSum3Backtrack(&res, []int{}, k, n, 1)
	return res
}

func combinationSum3Backtrack(res *[][]int, current []int, k int, n int, i int) {
	if n == 0 && len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*res = append(*res, temp)
		return
	}

	for j := i; j <= 9; j++ {
		current = append(current, j)
		combinationSum3Backtrack(res, current, k, n-j, j+1)
		current = current[:len(current)-1]
	}
}

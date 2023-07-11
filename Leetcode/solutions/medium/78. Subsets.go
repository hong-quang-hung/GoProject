package medium

import "fmt"

// Reference: https://leetcode.com/problems/subsets/
func init() {
	Solutions[78] = func() {
		fmt.Println("Input: nums = [1,2,3]")
		fmt.Println("Output:", subsets([]int{1, 2, 3}))
		fmt.Println("Input: nums = [2,2,2]")
		fmt.Println("Output:", subsets([]int{2, 2, 2}))
	}
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	subsetsBacktack(nums, &res, []int{}, 0)
	return res
}

func subsetsBacktack(nums []int, res *[][]int, current []int, i int) {
	dest := make([]int, len(current))
	copy(dest, current)
	*res = append(*res, dest)

	for j := i; j < len(nums); j++ {
		current = append(current, nums[j])
		subsetsBacktack(nums, res, current, j+1)
		current = current[:len(current)-1]
	}
}

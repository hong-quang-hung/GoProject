package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/next-permutation/
func init() {
	Solutions[31] = func() {
		var nums []int
		fmt.Println(`Input: nums = [1,2,3]`)
		nums = []int{1, 2, 3}
		nextPermutation(nums)
		fmt.Println(`Output:`, nums)
		fmt.Println(`Input: nums = [3,2,1]`)
		nums = []int{3, 2, 1}
		nextPermutation(nums)
		fmt.Println(`Output:`, nums)
		fmt.Println(`Input: nums = [1,3,2]`)
		nums = []int{1, 3, 2}
		nextPermutation(nums)
		fmt.Println(`Output:`, nums)
	}
}

func nextPermutation(nums []int) {
	j := len(nums) - 1
	for j > 0 && nums[j-1] >= nums[j] {
		j--
	}

	if j != 0 {
		l := len(nums) - 1
		for l > j-1 && nums[j-1] >= nums[l] {
			l--
		}
		nums[j-1], nums[l] = nums[l], nums[j-1]
	}

	for k := len(nums) - 1; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]
	}
}

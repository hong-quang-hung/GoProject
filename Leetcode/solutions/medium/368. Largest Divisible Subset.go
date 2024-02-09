package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/largest-divisible-subset/
func init() {
	Solutions[368] = func() {
		fmt.Println(`Input: nums = [1,2,3]`)
		fmt.Println(`Output:`, largestDivisibleSubset([]int{1, 2, 3}))
		fmt.Println(`Input: nums = [1,2,4,8]`)
		fmt.Println(`Output:`, largestDivisibleSubset([]int{1, 2, 4, 8}))
	}
}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	res := []int{}
	subsets := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cur := []int{}
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if len(subsets[j]) > len(cur) {
					cur = subsets[j]
				}
			}
		}

		newSubset := make([]int, len(cur)+1)
		copy(newSubset, cur)
		newSubset[len(cur)] = nums[i]
		subsets[i] = newSubset

		if len(subsets[i]) > len(res) {
			res = subsets[i]
		}
	}
	return res
}

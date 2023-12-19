package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/majority-element-ii/
func init() {
	Solutions[229] = func() {
		fmt.Println(`Input: nums = [3,2,3]`)
		fmt.Println(`Output:`, majorityElement([]int{3, 2, 3}))
		fmt.Println(`Input: nums = [1]`)
		fmt.Println(`Output:`, majorityElement([]int{1}))
		fmt.Println(`Input: nums = [1,2]`)
		fmt.Println(`Output:`, majorityElement([]int{1, 2}))
	}
}

func majorityElement(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)
	res := []int{}
	for i := 0; i < n; {
		j := 0
		for i+j < n && nums[i] == nums[i+j] {
			j++
		}
		if j > n/3 {
			res = append(res, nums[i])
		}
		i += j
	}
	return res
}

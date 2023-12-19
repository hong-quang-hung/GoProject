package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/neither-minimum-nor-maximum/
func init() {
	Solutions[2733] = func() {
		fmt.Println(`Input: nums = [3,2,1,4]`)
		fmt.Println(`Output:`, findNonMinOrMax([]int{3, 2, 1, 4}))
		fmt.Println(`Input: nums = [2,1,3]`)
		fmt.Println(`Output:`, findNonMinOrMax([]int{2, 1, 3}))
	}
}

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	sort.Ints(nums)
	return nums[1]
}

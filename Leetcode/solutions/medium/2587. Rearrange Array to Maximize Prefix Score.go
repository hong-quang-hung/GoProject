package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/rearrange-array-to-maximize-prefix-score/
func init() {
	Solutions[2587] = func() {
		fmt.Println(`Input: nums = [2,-1,0,1,-3,3,-3]`)
		fmt.Println(`Output:`, maxScore([]int{2, -1, 0, 1, -3, 3, -3}))
	}
}

func maxScore(nums []int) int {
	res, prefix := 0, int64(0)
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for _, num := range nums {
		prefix += int64(num)
		if prefix > 0 {
			res++
		} else {
			break
		}
	}
	return res
}

package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/longest-increasing-subsequence/
func init() {
	Solutions[300] = func() {
		fmt.Println(`Input: nums = [10,9,2,5,3,7,101,18]`)
		fmt.Println(`Output:`, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
		fmt.Println(`Input: nums = [0,1,0,3,2,3]`)
		fmt.Println(`Output:`, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
		fmt.Println(`Input: nums = [7,7,7,7,7,7,7]`)
		fmt.Println(`Output:`, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	}
}

func lengthOfLIS(nums []int) int {
	res := []int{}
	for _, num := range nums {
		if len(res) == 0 || res[len(res)-1] < num {
			res = append(res, num)
		} else {
			idx := sort.SearchInts(res, num)
			res[idx] = num
		}
	}
	return len(res)
}

package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/
func init() {
	Solutions[2966] = func() {
		fmt.Println(`Input: nums = [1,3,4,8,7,9,3,5,1], k = 2`)
		fmt.Println(`Output:`, divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
		fmt.Println(`Input: nums = [1,3,3,2,7,3], k = 3`)
		fmt.Println(`Output:`, divideArray([]int{1, 3, 3, 2, 7, 3}, 3))
	}
}

func divideArray(nums []int, k int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < n; i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		} else {
			res = append(res, nums[i:i+3])
		}
	}
	return res
}

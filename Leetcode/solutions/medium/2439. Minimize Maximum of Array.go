package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimize-maximum-of-array/
func init() {
	Solutions[2439] = func() {
		fmt.Println(`Input: nums = [3,7,1,6]`)
		fmt.Println(`Output:`, minimizeArrayValue([]int{3, 7, 1, 6}))
		fmt.Println(`Input: nums = [13,13,20,0,8,9,9]`)
		fmt.Println(`Output:`, minimizeArrayValue([]int{13, 13, 20, 0, 8, 9, 9}))
	}
}

func minimizeArrayValue(nums []int) int {
	sum, res := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum%(i+1) == 0 {
			res = max(res, sum/(i+1))
		} else {
			res = max(res, sum/(i+1)+1)
		}
	}
	return res
}

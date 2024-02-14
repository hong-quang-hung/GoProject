package medium

import "fmt"

// Reference: https://leetcode.com/problems/rearrange-array-elements-by-sign/
func init() {
	Solutions[2149] = func() {
		fmt.Println(`Input: nums = [3,1,-2,-5,2,-4]`)
		fmt.Println(`Output:`, rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
		fmt.Println(`Input: nums = [-1,1]`)
		fmt.Println(`Output:`, rearrangeArray([]int{-1, 1}))
	}
}

func rearrangeArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	ip, in := 0, 0
	for i := 0; i < n; i += 2 {
		for in < n && nums[in] > 0 {
			in++
		}

		for ip < n && nums[ip] < 0 {
			ip++
		}

		res[i], res[i+1] = nums[ip], nums[in]
		ip++
		in++
	}
	return res
}

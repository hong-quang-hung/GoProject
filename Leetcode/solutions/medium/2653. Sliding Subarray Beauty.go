package medium

import "fmt"

// Reference: https://leetcode.com/problems/sliding-subarray-beauty/
func init() {
	Solutions[2653] = func() {
		fmt.Println(`Input: nums = [1,-1,-3,-2,3], k = 3, x = 2`)
		fmt.Println(`Output:`, getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
		fmt.Println(`Input: nums = [-1,-2,-3,-4,-5], k = 2, x = 2`)
		fmt.Println(`Output:`, getSubarrayBeauty([]int{-1, -2, -3, -4, -5}, 2, 2))
		fmt.Println(`Input: nums = [-3,1,2,-3,0,-3], k = 2, x = 1`)
		fmt.Println(`Output:`, getSubarrayBeauty([]int{-3, 1, 2, -3, 0, -3}, 2, 1))
	}
}

func getSubarrayBeauty(nums []int, k int, x int) []int {
	n := len(nums)
	counter := make([]int, 50)
	res := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			counter[nums[i]+50]++
		}

		if i-k >= 0 && nums[i-k] < 0 {
			counter[nums[i-k]+50]--
		}

		if i-k+1 < 0 {
			continue
		}

		count := 0
		for j := 0; j < 50; j++ {
			count += counter[j]
			if count >= x {
				res[i-k+1] = j - 50
				break
			}
		}
	}
	return res
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/subarray-sum-equals-k/
func init() {
	Solutions[560] = func() {
		fmt.Println("Input: nums = [1,1,1], k = 2")
		fmt.Println("Output:", subarraySum([]int{1, 1, 1}, 2))
		fmt.Println("Input: nums = [1,2,3], k = 3")
		fmt.Println("Output:", subarraySum([]int{1, 2, 3}, 1))
	}
}

func subarraySum(nums []int, k int) int {
	n := len(nums)
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	fmt.Println(prefix)
	return 0
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/subarray-sum-equals-k/
func init() {
	Solutions[560] = func() {
		fmt.Println("Input: nums = [1,1,1], k = 2")
		fmt.Println("Output:", subarraySum([]int{1, 1, 1}, 2))
		fmt.Println("Input: nums = [1,2,3], k = 3")
		fmt.Println("Output:", subarraySum([]int{1, 2, 3}, 3))
	}
}

func subarraySum(nums []int, k int) int {
	m := make(map[int]int, len(nums))
	m[0]++
	res, sum := 0, 0
	for _, num := range nums {
		sum += num
		res += m[sum-k]
		m[sum]++
	}
	return res
}

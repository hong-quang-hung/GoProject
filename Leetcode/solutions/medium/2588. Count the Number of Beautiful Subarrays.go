package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-the-number-of-beautiful-subarrays/
func init() {
	Solutions[2588] = func() {
		fmt.Println("Input: nums = [4,3,1,2,4]")
		fmt.Println("Output:", beautifulSubarrays([]int{4, 3, 1, 2, 4}))
	}
}

func beautifulSubarrays(nums []int) int64 {
	var res int64 = 0
	var xor int64 = 0

	m := make(map[int64]int)
	m[0] = 1
	for _, num := range nums {
		xor ^= int64(num)
		x, v := m[xor]
		if v {
			res += int64(x)
		}
		m[xor] = x + 1
	}
	return res
}

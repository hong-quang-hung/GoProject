package medium

import "fmt"

// Reference: https://leetcode.com/problems/binary-subarrays-with-sum/
func init() {
	Solutions[930] = func() {
		fmt.Println(`Input: nums = [1,0,1,0,1], goal = 2`)
		fmt.Println(`Output:`, numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
		fmt.Println(`Input: nums = [0,0,0,0,0], goal = 0`)
		fmt.Println(`Output:`, numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
	}
}

func numSubarraysWithSum(nums []int, goal int) int {
	m := make(map[int]int, len(nums))
	m[0]++
	res, sum := 0, 0
	for _, num := range nums {
		sum += num
		res += m[sum-goal]
		m[sum]++
	}
	return res
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/jump-game/
func init() {
	Solutions[55] = func() {
		fmt.Println(`Input: nums = [2,3,1,1,4]`)
		fmt.Println(`Output:`, canJump([]int{2, 3, 1, 1, 4}))
		fmt.Println(`Input: nums = [3,2,1,0,4]`)
		fmt.Println(`Output:`, canJump([]int{3, 2, 1, 0, 4}))
	}
}

func canJump(nums []int) bool {
	n := len(nums)
	lastJump := n - 1
	for i := n - 2; i >= 0; i-- {
		if lastJump <= nums[i]+i {
			lastJump = i
		}
	}
	return lastJump == 0
}

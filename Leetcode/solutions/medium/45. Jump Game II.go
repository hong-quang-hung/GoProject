package medium

import "fmt"

// Reference: https://leetcode.com/problems/jump-game-ii/
func Leetcode_Jumb_Game_II() {
	fmt.Println("Input: nums = [2,3,1,1,4]")
	fmt.Println("Output:", jump_ii([]int{2, 3, 1, 1, 4}))
	fmt.Println("Input: nums = [2,3,0,1,4]]")
	fmt.Println("Output:", jump_ii([]int{2, 3, 0, 1, 4}))
}

func jump_ii(nums []int) int {
	res, last, next, n := 0, 0, 0, len(nums)
	for i := 0; i < n-1; i++ {
		next = max(next, i+nums[i])
		if next >= n-1 {
			res++
			break
		}
		if i == last {
			res++
			last = next
		}
	}
	return res
}

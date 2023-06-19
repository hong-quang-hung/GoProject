package easy

import "fmt"

// Reference: https://leetcode.com/problems/move-zeroes/
func Leetcode_Move_Zeroes() {
	fmt.Println("Input: nums = [0,1,0,3,12]")
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes(nums1)
	fmt.Println("Output:", nums1)
}

func moveZeroes(nums []int) {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

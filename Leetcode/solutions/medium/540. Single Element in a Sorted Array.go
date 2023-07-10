package medium

import "fmt"

func init() {
	Solutions[540] = Leetcode_Single_Non_Duplicate
}

// Reference: https://leetcode.com/problems/single-element-in-a-sorted-array/
func Leetcode_Single_Non_Duplicate() {
	fmt.Println("Input: nums = [1,1,2,3,3,4,4,8,8]")
	fmt.Println("Output:", singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println("Input: nums = [3,3,7,7,10,11,11]")
	fmt.Println("Output:", singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func singleNonDuplicate(nums []int) int {
	s, e := 0, len(nums)-1
	for s < e {
		if nums[s] != nums[s+1] {
			return nums[s]
		}

		if nums[e] != nums[e-1] {
			return nums[e]
		}
		s += 2
		e -= 2
	}
	return nums[s]
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func Leetcode_Remove_Duplicates() {
	fmt.Println("Input: nums = [1,1,1,2,2,3]")
	fmt.Println("Output:", removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println("Input: nums = [0,0,1,1,1,1,2,3,3]")
	fmt.Println("Output:", removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
}

func removeDuplicates(nums []int) int {
	i, c := 0, false
	for j := 1; j < len(nums); j++ {
		check := nums[i] == nums[j]
		if !check || !c {
			i++
			nums[i] = nums[j]
		}
		c = check
	}
	return i + 1
}

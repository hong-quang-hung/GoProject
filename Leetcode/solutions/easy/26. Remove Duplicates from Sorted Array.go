package easy

import "fmt"

func init() {
	Solutions[26] = Leetcode_Remove_Duplicates
}

// Reference: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func Leetcode_Remove_Duplicates() {
	fmt.Println("Input: nums = [0,0,1,1,1,2,2,3,3,4]")
	fmt.Println("Output:", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

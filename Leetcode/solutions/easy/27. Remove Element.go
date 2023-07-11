package easy

import "fmt"

// Reference: https://leetcode.com/problems/remove-element/
func init() {
	Solutions[27] = func() {
		fmt.Println("Input: nums = [3,2,2,3], val = 3")
		fmt.Println("Output:", removeElement([]int{3, 2, 2, 3}, 3))
	}
}

func removeElement(nums []int, val int) int {
	var res int = 0
	for _, num := range nums {
		if num != val {
			nums[res] = num
			res++
		}
	}
	return res
}

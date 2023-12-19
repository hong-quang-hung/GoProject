package easy

import "fmt"

// Reference: https://leetcode.com/problems/majority-element/
func init() {
	Solutions[169] = func() {
		fmt.Println(`Input: nums = [3,2,3]`)
		fmt.Println(`Output:`, majorityElement([]int{3, 2, 3}))
		fmt.Println(`Input: nums = [2,2,1,1,1,2,2]`)
		fmt.Println(`Output:`, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	}
}

func majorityElement(nums []int) int {
	res := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
		}

		if res == nums[i] {
			count++
		} else {
			count--
		}
	}
	return res
}

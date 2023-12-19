package easy

import "fmt"

// Reference: https://leetcode.com/problems/sort-array-by-parity/
func init() {
	Solutions[905] = func() {
		fmt.Println(`Input: nums = [3,1,2,4]`)
		fmt.Println(`Output:`, sortArrayByParity([]int{3, 1, 2, 4}))
		fmt.Println(`Input: nums = [0]`)
		fmt.Println(`Output:`, sortArrayByParity([]int{0}))
	}
}

func sortArrayByParity(nums []int) []int {
	for left, right := 0, len(nums)-1; left <= right; {
		if nums[left]&1 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return nums
}

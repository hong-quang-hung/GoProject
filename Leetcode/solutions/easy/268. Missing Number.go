package easy

import "fmt"

// Reference: https://leetcode.com/problems/missing-number/
func init() {
	Solutions[268] = func() {
		fmt.Println(`Input: nums = [3,0,1]`)
		fmt.Println(`Output:`, missingNumber([]int{3, 0, 1}))
		fmt.Println(`Input: nums = [0,1]`)
		fmt.Println(`Output:`, missingNumber([]int{0, 1}))
		fmt.Println(`Input: nums = [9,6,4,2,3,5,7,0,1]`)
		fmt.Println(`Output:`, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	}
}

func missingNumber(nums []int) int {
	res := len(nums)
	for i, num := range nums {
		res += i - num
	}
	return res
}

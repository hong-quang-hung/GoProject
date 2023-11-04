package easy

import "fmt"

// Reference: https://leetcode.com/problems/plus-one/
func init() {
	Solutions[66] = func() {
		fmt.Println("Input: digits = [1,2,3]")
		fmt.Println("Output:", plusOne([]int{1, 2, 3}))
		fmt.Println("Input: digits = [4,3,2,1]")
		fmt.Println("Output:", plusOne([]int{4, 3, 2, 1}))
		fmt.Println("Input: digits = [9]")
		fmt.Println("Output:", plusOne([]int{9}))
	}
}

func plusOne(digits []int) []int {
	plus, idx := 1, len(digits)-1
	for plus > 0 && idx >= 0 {
		plus, digits[idx], idx = (plus+digits[idx])/10, (plus+digits[idx])%10, idx-1
	}

	if plus > 0 {
		digits = append([]int{plus}, digits...)
	}
	return digits
}

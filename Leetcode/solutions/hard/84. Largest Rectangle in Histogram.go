package hard

import "fmt"

// Reference: https://leetcode.com/problems/largest-rectangle-in-histogram/
func init() {
	Solutions[84] = func() {
		fmt.Println("Input: heights = [2,1,5,6,2,3]")
		fmt.Println("Output:", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
		fmt.Println("Input: heights = [2,4]")
		fmt.Println("Output:", largestRectangleArea([]int{2, 4}))
	}
}

func largestRectangleArea(heights []int) int {
	res := 0
	return res
}

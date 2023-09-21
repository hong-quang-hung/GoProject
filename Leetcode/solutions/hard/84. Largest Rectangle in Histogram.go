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
	stack := []int{}
	arr := make([]int, len(heights))
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] < h {
			pop := stack[len(stack)-1]
			arr[pop] = i - pop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	fmt.Println(arr)
	return res
}

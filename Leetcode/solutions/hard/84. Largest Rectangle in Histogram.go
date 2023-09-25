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
	stack, res := [][2]int{}, 0
	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			idx, v := stack[len(stack)-1][0], stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			res = max(res, (i-idx)*v)
			start = idx
		}
		stack = append(stack, [2]int{start, h})
	}

	for i := 0; i < len(stack); i++ {
		idx, h := stack[i][0], stack[i][1]
		res = max(res, h*(len(heights)-idx))
	}
	return res
}

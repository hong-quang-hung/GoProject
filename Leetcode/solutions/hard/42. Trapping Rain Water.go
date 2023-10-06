package hard

import "fmt"

// Reference: https://leetcode.com/problems/trapping-rain-water/
func init() {
	Solutions[42] = func() {
		fmt.Println("Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]")
		fmt.Println("Output:", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
		fmt.Println("Input: height = [4,2,0,3,2,5]")
		fmt.Println("Output:", trap([]int{4, 2, 0, 3, 2, 5}))
		fmt.Println("Input: height = [4,2,3]")
		fmt.Println("Output:", trap([]int{4, 2, 3}))
	}
}

func trap(height []int) int {
	n := len(height)
	stack := []int{}
	m := make([]int, n)
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= height[i] {
			pop := stack[len(stack)-1]
			m[pop] = i - pop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	i := 0
	for i < n {
		if m[i] == 0 {
			break
		}

		for j := 1; j < m[i]; j++ {
			res += height[i] - height[i+j]
		}
		i += m[i]
	}

	fmt.Println(i, m)
	return res
}

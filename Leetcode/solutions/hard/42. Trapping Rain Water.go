package hard

import "fmt"

// Reference: https://leetcode.com/problems/trapping-rain-water/
func init() {
	Solutions[42] = func() {
		// fmt.Println("Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]")
		// fmt.Println("Output:", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
		// fmt.Println("Input: height = [4,2,0,3,2,5]")
		// fmt.Println("Output:", trap([]int{4, 2, 0, 3, 2, 5}))
		fmt.Println("Input: height = [4,2,3]")
		fmt.Println("Output:", trap([]int{4, 2, 3}))
	}
}

func trap(height []int) int {
	n := len(height)
	stack := []int{}
	m := make([]int, n)
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= h {
			pop := stack[len(stack)-1]
			m[pop] = i - pop
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	fmt.Println(m)

	res := 0
	cur := 0
	for i := 0; i < n; {
		if m[i] == 0 {
			i++
		} else {
			for j := 1; j < m[i]; j++ {
				res += height[cur] - height[i+j]
			}
			cur += m[i]
			i += m[i]
		}
	}
	return res
}

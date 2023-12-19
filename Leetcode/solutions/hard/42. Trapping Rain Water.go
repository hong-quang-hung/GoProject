package hard

import "fmt"

// Reference: https://leetcode.com/problems/trapping-rain-water/
func init() {
	Solutions[42] = func() {
		fmt.Println(`Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]`)
		fmt.Println(`Output:`, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
		fmt.Println(`Input: height = [4,2,0,3,2,5]`)
		fmt.Println(`Output:`, trap([]int{4, 2, 0, 3, 2, 5}))
		fmt.Println(`Input: height = [4,2,3]`)
		fmt.Println(`Output:`, trap([]int{4, 2, 3}))
	}
}

func trap(height []int) int {
	n := len(height)
	stack1, stack2 := []int{}, []int{}
	m1, m2 := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack1) > 0 && height[stack1[len(stack1)-1]] <= height[i] {
			pop := stack1[len(stack1)-1]
			m1[pop] = i - pop
			stack1 = stack1[:len(stack1)-1]
		}
		stack1 = append(stack1, i)
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack2) > 0 && height[stack2[len(stack2)-1]] <= height[i] {
			pop := stack2[len(stack2)-1]
			m2[pop] = pop - i
			stack2 = stack2[:len(stack2)-1]
		}
		stack2 = append(stack2, i)
	}

	res := 0
	left := 0
	for left < n {
		if m1[left] == 0 {
			break
		}
		for j := 1; j < m1[left]; j++ {
			res += height[left] - height[left+j]
		}
		left += m1[left]
	}

	for i := n - 1; i > left; {
		if m2[i] == 0 {
			break
		}
		for j := 1; j < m2[i]; j++ {
			res += height[i] - height[i-j]
		}
		i -= m2[i]
	}
	return res
}

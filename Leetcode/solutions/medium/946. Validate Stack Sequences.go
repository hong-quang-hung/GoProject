package medium

import "fmt"

// Reference: https://leetcode.com/problems/validate-stack-sequences/
func init() {
	Solutions[946] = func() {
		fmt.Println(`Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]`)
		fmt.Println(`Output:`, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
		fmt.Println(`Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]`)
		fmt.Println(`Output:`, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
	}
}

func validateStackSequences(pushed []int, popped []int) bool {
	j, n := 0, len(pushed)
	stack := []int{}
	for _, p := range pushed {
		stack = append(stack, p)
		for len(stack) > 0 && j < n && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
	}
	return j == n
}

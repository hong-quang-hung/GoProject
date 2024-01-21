package medium

import "fmt"

// Reference: https://leetcode.com/problems/sum-of-subarray-minimums/
func init() {
	Solutions[907] = func() {
		fmt.Println(`Input: arr = [3,1,2,4]`)
		fmt.Println(`Output:`, sumSubarrayMins([]int{3, 1, 2, 4}))
		fmt.Println(`Input: arr = [11,81,94,43,3]`)
		fmt.Println(`Output:`, sumSubarrayMins([]int{11, 81, 94, 43, 3}))
	}
}

func sumSubarrayMins(arr []int) int {
	stack := [][]int{{-1, 0}}
	res := 0
	for i := 0; i < len(arr); i++ {
		for stack[len(stack)-1][0] >= 0 && arr[stack[len(stack)-1][0]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}

		x := stack[len(stack)-1]
		v := arr[i]*(i-x[0]) + x[1]
		stack = append(stack, []int{i, v})
		res += v
		res = res % mod
	}
	return res
}

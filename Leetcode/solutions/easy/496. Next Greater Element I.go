package easy

import "fmt"

// Reference: https://leetcode.com/problems/next-greater-element-i/
func init() {
	Solutions[496] = func() {
		fmt.Println(`Input: nums1 = [4,1,2], nums2 = [1,3,4,2]`)
		fmt.Println(`Output:`, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
		fmt.Println(`Input: nums1 = [2,4], nums2 = [1,2,3,4]`)
		fmt.Println(`Output:`, nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	}
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	next := make([]int, 10001)
	for _, num := range nums2 {
		next[num] = -1
		for len(stack) > 0 && stack[len(stack)-1] < num {
			next[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	res := []int{}
	for _, num := range nums1 {
		res = append(res, next[num])
	}
	return res
}

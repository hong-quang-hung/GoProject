package medium

import "fmt"

// Reference: https://leetcode.com/problems/container-with-most-water/
func init() {
	Solutions[11] = func() {
		fmt.Println(`Input: height = [1,8,6,2,5,4,8,3,7]`)
		fmt.Println(`Output:`, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
		fmt.Println(`Input: height = [1,1]`)
		fmt.Println(`Output:`, maxArea([]int{1, 1}))
	}
}

func maxArea(height []int) int {
	res, n := 0, len(height)
	left, right := 0, n-1
	for left < right {
		res = max(res, min(height[left], height[right])*(right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}

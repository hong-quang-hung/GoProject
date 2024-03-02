package easy

import "fmt"

// Reference: https://leetcode.com/problems/squares-of-a-sorted-array/
func init() {
	Solutions[977] = func() {
		fmt.Println(`Input: nums = [-4,-1,0,3,10]`)
		fmt.Println(`Output:`, sortedSquares([]int{-4, -1, 0, 3, 10}))
		fmt.Println(`Input: nums = [-7,-3,2,3,11]`)
		fmt.Println(`Output:`, sortedSquares([]int{-7, -3, 2, 3, 11}))
	}
}

func sortedSquares(nums []int) []int {
	m := make([]int, 10000)
	for _, num := range nums {
		m[abs(num)]++
	}

	j := 0
	for i, num := range m {
		for z := 0; z < num; z++ {
			nums[j+z] = i * i
		}
		j += num
	}
	return nums
}

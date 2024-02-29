package easy

import "fmt"

// Reference: https://leetcode.com/problems/split-the-array/
func init() {
	Solutions[3046] = func() {
		fmt.Println(`Input: nums = [1,1,2,2,3,4]`)
		fmt.Println(`Output:`, isPossibleToSplit([]int{1, 1, 2, 2, 3, 4}))
		fmt.Println(`Input: nums = [1,1,1,1]`)
		fmt.Println(`Output:`, isPossibleToSplit([]int{1, 1, 1, 1}))
	}
}

func isPossibleToSplit(nums []int) bool {
	m := make([]int, 101)
	for _, num := range nums {
		if m[num] == 2 {
			return false
		}
		m[num]++
	}
	return true
}

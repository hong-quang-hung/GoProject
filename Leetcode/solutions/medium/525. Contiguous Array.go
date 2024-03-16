package medium

import "fmt"

// Reference: https://leetcode.com/problems/contiguous-array/
func init() {
	Solutions[525] = func() {
		fmt.Println(`Input: nums = [0,1]`)
		fmt.Println(`Output:`, findMaxLength([]int{0, 1}))
		fmt.Println(`Input: nums = [0,1,0]`)
		fmt.Println(`Output:`, findMaxLength([]int{0, 1, 0}))
	}
}

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1
	res, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}

		if v, ok := m[count]; ok {
			res = max(res, i-v)
		} else {
			m[count] = i
		}
	}
	return res
}

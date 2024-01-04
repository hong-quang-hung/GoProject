package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
func init() {
	Solutions[2870] = func() {
		fmt.Println(`Input: nums = [2,3,3,2,2,4,2,3,4]`)
		fmt.Println(`Output:`, minOperations2([]int{2, 3, 3, 2, 2, 4, 2, 3, 4}))
		fmt.Println(`Input: nums = [2,1,2,2,3,3]`)
		fmt.Println(`Output:`, minOperations2([]int{2, 1, 2, 2, 3, 3}))
	}
}

func minOperations2(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	res := 0
	for _, val := range m {
		if val == 1 {
			return -1
		} else {
			res += val / 3
			if val%3 > 0 {
				res++
			}
		}
	}
	return res
}

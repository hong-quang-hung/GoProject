package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-nice-pairs-in-an-array/
func init() {
	Solutions[1814] = func() {
		fmt.Println(`Input: nums = [42,11,1,97]`)
		fmt.Println(`Output:`, countNicePairs([]int{42, 11, 1, 97}))
		fmt.Println(`Input: nums = [13,10,35,24,76]`)
		fmt.Println(`Output:`, countNicePairs([]int{13, 10, 35, 24, 76}))
	}
}

func countNicePairs(nums []int) int {
	rev := func(i int) int {
		r := 0
		for i > 0 {
			r = r*10 + i%10
			i /= 10
		}
		return r
	}

	res := 0
	m := make(map[int]int)
	for _, num := range nums {
		key := num - rev(num)
		res = (res + m[key]) % mod
		m[key]++
	}
	return res
}

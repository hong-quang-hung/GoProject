package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-consecutive-sequence/
func init() {
	Solutions[128] = func() {
		fmt.Println(`Input: nums = [100,4,200,1,3,2]`)
		fmt.Println(`Output:`, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
		fmt.Println(`Input: nums = [0,3,7,2,5,8,4,6,0,1]`)
		fmt.Println(`Output:`, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	}
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	res := 0
	for n := range m {
		if _, ok := m[n-1]; !ok {
			l := 1
			for {
				if _, ok = m[n+l]; ok {
					l++
					continue
				}

				res = max(l, res)
				break
			}
		}
	}
	return res
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/
func init() {
	Solutions[1503] = func() {
		fmt.Println("Input: n = 4, left = [4,3], right = [0,1]")
		fmt.Println("Output:", getLastMoment(4, []int{4, 3}, []int{0, 1}))
		fmt.Println("Input: n = 7, left = [], right = [0,1,2,3,4,5,6,7]")
		fmt.Println("Output:", getLastMoment(7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}))
		fmt.Println("Input: n = 7, left = [0,1,2,3,4,5,6,7], right = []")
		fmt.Println("Output:", getLastMoment(7, []int{0, 1, 2, 3, 4, 5, 6, 7}, []int{}))
	}
}

func getLastMoment(n int, left []int, right []int) int {
	res := 0
	for _, a := range left {
		res = max(res, a)
	}
	for _, a := range right {
		res = max(res, n-a)
	}
	return res
}

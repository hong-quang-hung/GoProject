package medium

import "fmt"

// Reference: https://leetcode.com/problems/build-an-array-with-stack-operations/
func init() {
	Solutions[1441] = func() {
		fmt.Println(`Input: target = [1,3], n = 3`)
		fmt.Println(`Output:`, buildArray([]int{1, 3}, 3))
		fmt.Println(`Input: target = [1,2,3], n = 3`)
		fmt.Println(`Output:`, buildArray([]int{1, 2, 3}, 3))
		fmt.Println(`Input: target = [1,2], n = 4`)
		fmt.Println(`Output:`, buildArray([]int{1, 2}, 4))
	}
}

func buildArray(target []int, n int) []string {
	res := []string{}
	currVal := 1
	for _, t := range target {
		for i := currVal; i < t; i++ {
			res = append(res, `Push`, `Pop`)
		}
		res = append(res, `Push`)
		currVal = t + 1
	}
	return res
}

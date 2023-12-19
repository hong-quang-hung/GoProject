package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/constrained-subsequence-sum/
func init() {
	Solutions[1425] = func() {
		fmt.Println(`Input: nums = [10,2,-10,5,20], k = 2`)
		fmt.Println(`Output:`, constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
		fmt.Println(`Input: nums = [-1,-2,-3], k = 1`)
		fmt.Println(`Output:`, constrainedSubsetSum([]int{-1, -2, -3}, 1))
		fmt.Println(`Input: nums = [10,-2,-10,-5,20], k = 2`)
		fmt.Println(`Output:`, constrainedSubsetSum([]int{10, -2, -10, -5, 20}, 2))
	}
}

func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	q := [][]int{}
	res := math.MinInt
	for i := 0; i < n; i++ {
		for len(q) > 0 && i-q[0][1] > k {
			q = q[1:]
		}

		value := nums[i]
		if len(q) > 0 {
			value = max(value, value+q[0][0])
		}

		for len(q) > 0 && q[len(q)-1][0] < value {
			q = q[:len(q)-1]
		}

		if value > 0 {
			q = append(q, []int{value, i})
		}

		res = max(res, value)
	}
	return res
}

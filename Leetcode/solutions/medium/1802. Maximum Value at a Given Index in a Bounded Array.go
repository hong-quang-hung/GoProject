package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-value-at-a-given-index-in-a-bounded-array/
func init() {
	Solutions[1802] = func() {
		fmt.Println(`Input: n = 4, index = 2, maxSum = 6`)
		fmt.Println(`Output:`, maxValue(4, 2, 6))
		fmt.Println(`Input: n = 6, index = 1, maxSum = 10`)
		fmt.Println(`Output:`, maxValue(4, 2, 6))
	}
}

func maxValue(n int, index int, maxSum int) int {
	if n == maxSum {
		return 1
	}

	res := (maxSum + (index*(index+1)+(n-1-index)*(n-index))/2) / n
	if res > index && res > n-1-index {
		return res
	}

	total, left, right := n, 0, 0
	res = 1
	for total < maxSum {
		res++
		if left < index {
			left++
		}

		if right < n-1-index {
			right++
		}

		total += left + right + 1
	}
	return res
}

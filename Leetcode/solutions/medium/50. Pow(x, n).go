package medium

import "fmt"

// Reference: https://leetcode.com/problems/powx-n/
func init() {
	Solutions[50] = func() {
		fmt.Println("Input: x = 2.00000, n = 10")
		fmt.Println("Output:", myPow(2.00000, 10))
		fmt.Println("Input: x = 2.00000, n = -1")
		fmt.Println("Output:", myPow(2.00000, -1))
	}
}

func myPow(x float64, n int) float64 {
	res, m := float64(1), x
	if n < 0 {
		n = -n
		m = 1 / m
	}

	for n > 0 {
		if n%2 == 1 {
			res = res * m
			n = n - 1
		} else {
			m = m * m
			n = n / 2
		}
	}
	return res
}

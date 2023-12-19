package easy

import "fmt"

// Reference: https://leetcode.com/problems/ugly-number/
func init() {
	Solutions[263] = func() {
		fmt.Println(`Input: n = 6`)
		fmt.Println(`Output:`, isUgly(6))
		fmt.Println(`Input: n = 14`)
		fmt.Println(`Output:`, isUgly(14))
		fmt.Println(`Input: n = 40`)
		fmt.Println(`Output:`, isUgly(40))
	}
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 || n == 2 || n == 3 || n == 5 {
		return true
	}

	if n%2 == 0 {
		return isUgly(n / 2)
	}

	if n%3 == 0 {
		return isUgly(n / 3)
	}

	if n%5 == 0 {
		return isUgly(n / 5)
	}
	return false
}

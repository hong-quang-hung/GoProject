package easy

import "fmt"

// Reference: https://leetcode.com/problems/calculate-money-in-leetcode-bank/
func init() {
	Solutions[1716] = func() {
		fmt.Println("Input: n = 4")
		fmt.Println("Output:", totalMoney(4))
		fmt.Println("Input: n = 10")
		fmt.Println("Output:", totalMoney(10))
		fmt.Println("Input: n = 20")
		fmt.Println("Output:", totalMoney(20))
	}
}

func totalMoney(n int) int {
	f := func(u, i, d int) int {
		if i <= 0 {
			return 0
		} else if i <= 2 {
			return u + (i-1)*(u+d)
		} else {
			return i*u + i*(i-1)*d/2
		}
	}

	week := n / 7
	return f(week+1, n%7, 1) + f(28, week, 7)
}

package easy

import "fmt"

// Reference: https://leetcode.com/problems/sum-multiples/
func init() {
	Solutions[2652] = func() {
		fmt.Println("Input: n = 10")
		fmt.Println("Output:", sumOfMultiples(10))
		fmt.Println("Input: n = 9")
		fmt.Println("Output:", sumOfMultiples(9))
	}
}

func sumOfMultiples(n int) int {
	res := 0
	for i := 3; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}
	return res
}

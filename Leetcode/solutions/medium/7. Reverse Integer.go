package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/reverse-integer/
func Leetcode_Reverse() {
	fmt.Println("Input: x = 123")
	fmt.Println("Output:", reverse(123))
	fmt.Println("Input: x = -123")
	fmt.Println("Output:", reverse(-123))
	fmt.Println("Input: x = 120")
	fmt.Println("Output:", reverse(120))
}

func reverse(x int) int {
	stack := []int{}
	flag := false
	if x < 0 {
		flag = true
		x = -x
	}

	for x > 0 {
		stack = append([]int{x % 10}, stack...)
		x /= 10
	}

	res, ten := 0, 1
	for _, s := range stack {
		res = res + s*ten
		ten *= 10
	}

	if flag {
		res = -res
	}

	if res < -2147483648 || res > 2147483647 {
		return 0
	}
	return res
}

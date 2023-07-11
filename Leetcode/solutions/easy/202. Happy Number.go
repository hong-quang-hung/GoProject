package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/happy-number/
func init() {
	Solutions[202] = func() {
		fmt.Println("Input: num = 19")
		fmt.Println("Output:", isHappy(19))
		fmt.Println("Input: num = 2")
		fmt.Println("Output:", isHappy(2))
		fmt.Println("Input: num = 3")
		fmt.Println("Output:", isHappy(3))
	}
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		if m[n] {
			return false
		} else {
			m[n] = true
		}

		sum := 0
		for _, number := range fmt.Sprint(n) {
			sum += int(math.Pow(float64(number-'0'), 2))
		}
		n = sum
	}
	return true
}

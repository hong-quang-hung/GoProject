package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/count-the-digits-that-divide-a-number/
func Leetcode_Count_Digits() {
	fmt.Println("Input: num = 7")
	fmt.Println("Output:", countDigits(7))
	fmt.Println("Input: num = 121")
	fmt.Println("Output:", countDigits(121))
	fmt.Println("Input: num = 1248")
	fmt.Println("Output:", countDigits(1248))
}

func countDigits(num int) int {
	n, res := num, 0
	for n > 0 {
		val := n % 10
		if num%val == 0 {
			res++
		}
		n = (n - val) / 10
	}
	return res
}

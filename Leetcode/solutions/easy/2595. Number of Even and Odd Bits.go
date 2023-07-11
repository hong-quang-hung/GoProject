package easy

import "fmt"

// Reference: https://leetcode.com/problems/number-of-even-and-odd-bits/
func init() {
	Solutions[2595] = func() {
		fmt.Println("Input: n = 11")
		fmt.Println("Output:", evenOddBit(11))
	}
}

func evenOddBit(n int) []int {
	arr := []int{0, 0}
	var p int = 0
	for n > 0 {
		if n%2 == 1 {
			arr[p%2]++
		}
		p++
		n /= 2
	}
	return arr
}

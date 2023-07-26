package easy

import "fmt"

// Reference: https://leetcode.com/problems/guess-number-higher-or-lower/
func init() {
	Solutions[374] = func() {
		fmt.Println("Input: n = 10, pick = 6")
		fmt.Println("Output:", guessNumber(10))
		fmt.Println("Input: n = 1, pick = 1")
		fmt.Println("Output:", guessNumber(2))
		fmt.Println("Input: n = 2, pick = 1")
		fmt.Println("Output:", guessNumber(1))
	}
}
func guess(n int) int {
	return 0
}

func guessNumber(n int) int {
	left, right := 0, n
	for {
		mid := (left + right) / 2
		ans := guess(mid)
		if ans == 0 {
			return mid
		} else if ans == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

package hard

import "fmt"

// Reference: https://leetcode.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/
func init() {
	Solutions[1420] = func() {
		fmt.Println("Input: n = 2, m = 3, k = 1")
		fmt.Println("Output:", numOfArrays(2, 3, 1))
		fmt.Println("Input: n = 5, m = 2, k = 3")
		fmt.Println("Output:", numOfArrays(5, 2, 3))
		fmt.Println("Input: n = 9, m = 1, k = 1")
		fmt.Println("Output:", numOfArrays(9, 1, 1))
	}
}

func numOfArrays(n int, m int, k int) int {
	return 0
}

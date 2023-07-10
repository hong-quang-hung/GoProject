package easy

import "fmt"

func init() {
	Solutions[2652] = Leetcode_Sum_Of_Multiples
}

// Reference: https://leetcode.com/problems/sum-multiples/
func Leetcode_Sum_Of_Multiples() {
	fmt.Println("Input: n = 7")
	fmt.Println("Output:", sumOfMultiples(7))
	fmt.Println("Input: n = 10")
	fmt.Println("Output:", sumOfMultiples(10))
	fmt.Println("Input: n = 9")
	fmt.Println("Output:", sumOfMultiples(9))
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

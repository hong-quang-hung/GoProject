package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-the-divisibility-array-of-a-string/
func Leetcode_Divisibility_Array() {
	fmt.Println("Input: word = '998244353', m = 3")
	fmt.Println("Output:", divisibilityArray("998244353", 3))
}

func divisibilityArray(word string, m int) []int {
	res := make([]int, len(word))
	divisible := 0
	for i := 0; i < len(word); i++ {
		divisible = (divisible*10 + int(word[i]-'0')) % m
		if divisible == 0 {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}

package medium

import (
	"fmt"
	"strconv"
)

func init() {
	Solutions[2698] = Leetcode_Punishment_Number
}

// Reference: https://leetcode.com/problems/find-the-punishment-number-of-an-integer/
func Leetcode_Punishment_Number() {
	fmt.Println("Input: n = 10")
	fmt.Println("Output:", punishmentNumber(10))
	fmt.Println("Input: n = 37")
	fmt.Println("Output:", punishmentNumber(37))
}

func punishmentNumber(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		sqr := i * i
		if punishmentNumberCheck(strconv.Itoa(sqr), i) {
			res += sqr
		}
	}
	return res
}

func punishmentNumberCheck(s string, i int) bool {
	if len(s) == 0 && i == 0 {
		return true
	}

	if i < 0 {
		return false
	}

	for j := 0; j < len(s); j++ {
		left := s[0 : j+1]
		right := s[j+1:]
		leftNum, _ := strconv.Atoi(left)
		if punishmentNumberCheck(right, i-leftNum) {
			return true
		}
	}
	return false
}

package easy

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/maximum-odd-binary-number/
func init() {
	Solutions[2864] = func() {
		fmt.Println("Input: s = '0101'")
		fmt.Println("Output:", maximumOddBinaryNumber("0101"))
	}
}

func maximumOddBinaryNumber(s string) string {
	count := strings.Count(s, "1") - 1
	n := len(s)
	res := make([]byte, n)
	for i := 0; i < count; i++ {
		res[i] = '1'
	}
	for i := count; i < n-1; i++ {
		res[i] = '0'
	}
	res[n-1] = '1'
	return string(res)
}

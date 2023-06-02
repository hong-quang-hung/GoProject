package easy

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/remove-trailing-zeros-from-a-string/
func Leetcode_Remove_Trailing_Zeros() {
	fmt.Println("Input: num = '51230100'")
	fmt.Println("Output:", removeTrailingZeros("51230100"))
	fmt.Println("Input: num = '123'")
	fmt.Println("Output:", removeTrailingZeros("123"))
}

func removeTrailingZeros(num string) string {
	return strings.Trim(num, "0")
}

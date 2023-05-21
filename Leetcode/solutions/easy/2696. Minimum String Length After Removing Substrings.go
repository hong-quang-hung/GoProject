package easy

import (
	"fmt"
	"strings"
)

// Reference:
func Leetcode_Min_Length() {
	fmt.Println("Input: s = 'ABFCACDB'")
	fmt.Println("Output:", minLength("ABFCACDB"))
	fmt.Println("Input: s = 'ACBBD'")
	fmt.Println("Output:", minLength("ACBBD"))
}

func minLength(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}
	return len(s)
}

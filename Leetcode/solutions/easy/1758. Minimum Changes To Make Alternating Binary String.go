package easy

import "fmt"

// Reference: https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
func init() {
	Solutions[1758] = func() {
		fmt.Println(`Input: s = "0100"`)
		fmt.Println(`Output:`, minOperations("0100"))
		fmt.Println(`Input: s = "10"`)
		fmt.Println(`Output:`, minOperations("10"))
		fmt.Println(`Input: s = "1111"`)
		fmt.Println(`Output:`, minOperations("1111"))
	}
}

func minOperations(s string) int {
	zero, one := 0, 0
	for i, ch := range s {
		if i%2 == 0 && ch == '0' || i%2 == 1 && ch == '1' {
			one++
		} else {
			zero++
		}
	}
	return min(zero, one)
}

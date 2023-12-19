package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/excel-sheet-column-title/
func init() {
	Solutions[168] = func() {
		fmt.Println(`Input: columnNumber = 1`)
		fmt.Println(`Output:`, convertToTitle(1))
		fmt.Println(`Input: columnNumber = 28`)
		fmt.Println(`Output:`, convertToTitle(28))
		fmt.Println(`Input: columnNumber = 701`)
		fmt.Println(`Output:`, convertToTitle(701))
	}
}

func convertToTitle(columnNumber int) string {
	res := []byte{}
	for columnNumber >= 1 {
		columnNumber--
		res = append([]byte{byte(columnNumber%26 + int('A'))}, res...)
		columnNumber /= 26
	}
	return string(res)
}

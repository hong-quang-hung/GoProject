package medium

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/zigzag-conversion/
func init() {
	Solutions[6] = func() {
		fmt.Println("Input: x = 'PAYPALISHIRING', numRows = 3")
		fmt.Println("Output:", convert("PAYPALISHIRING", 3))
		fmt.Println("Input: x = 'PAYPALISHIRING', numRows = 4")
		fmt.Println("Output:", convert("PAYPALISHIRING", 4))
	}
}

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}

	zigzag := make([][]byte, numRows)
	res := strings.Builder{}
	for i := 0; i < numRows; i++ {
		plus := 2 * (numRows - 1)
		for j := i; j < n; j += plus {
			zigzag[i] = append(zigzag[i], s[j])
			if i != 0 && i != numRows-1 && j+(numRows-1-i)*2 < n {
				zigzag[i] = append(zigzag[i], s[j+(numRows-1-i)*2])
			}
		}
		res.Write(zigzag[i])
	}
	return res.String()
}

package medium

import "fmt"

func init() {
	Solutions[6] = Leetcode_Convert
}

// Reference: https://leetcode.com/problems/zigzag-conversion/
func Leetcode_Convert() {
	fmt.Println("Input: x = 'PAYPALISHIRING', numRows = 3")
	fmt.Println("Output:", convert("PAYPALISHIRING", 3))
	fmt.Println("Input: x = 'PAYPALISHIRING', numRows = 4")
	fmt.Println("Output:", convert("PAYPALISHIRING", 4))
	fmt.Println("Input: x = 'A', numRows = 1")
	fmt.Println("Output:", convert("A", 1))
}

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}

	zigzag := make([][]byte, numRows)
	res := ""
	for i := 0; i < numRows; i++ {
		plus := 2 * (numRows - 1)
		for j := i; j < n; j += plus {
			zigzag[i] = append(zigzag[i], s[j])
			if i != 0 && i != numRows-1 && j+(numRows-1-i)*2 < n {
				zigzag[i] = append(zigzag[i], s[j+(numRows-1-i)*2])
			}
		}
		res += string(zigzag[i])
	}
	return res
}

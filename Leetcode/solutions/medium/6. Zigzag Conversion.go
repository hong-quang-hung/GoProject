package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/zigzag-conversion/
func Leetcode_Convert() {
	fmt.Println("Input: x = 'PAYPALISHIRING', numRows = 3")
	fmt.Println("Output:", convert("PAYPALISHIRING", 3))
	fmt.Println("Input: x = 'PAYPALISHIRING', numRows = 4")
	fmt.Println("Output:", convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {
	n := len(s)
	g := int(math.Ceil(float64(n) / float64(numRows+1)))
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		if i%(numRows+1) == 0 {
			res[i/(numRows+1)] = s[i]
		} else {
			res[i/(numRows+1)*g+i%(numRows+1)] = s[i]
		}
	}
	return string(res)
}

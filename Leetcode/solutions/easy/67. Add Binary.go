package easy

import "fmt"

// Reference: https://leetcode.com/problems/add-binary/
func Leetcode_Add_Binary() {
	fmt.Println("Input: a = '1010', b = '1011'")
	fmt.Println("Output:", addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	s, c := "", false
	for len(a) > 0 || len(b) > 0 {
		bitA := len(a) > 0 && a[len(a)-1] == '1'
		bitB := len(b) > 0 && b[len(b)-1] == '1'

		if xor := c != bitA; xor != bitB {
			s = "1" + s
		} else {
			s = "0" + s
		}
		c = (bitA && bitB) || (c && (bitA || bitB))
		if len(a) > 0 {
			a = a[0 : len(a)-1]
		}
		if len(b) > 0 {
			b = b[0 : len(b)-1]
		}
	}
	if c {
		return "1" + s
	}
	return s
}

package easy

import "fmt"

func init() {
	Solutions[13] = Leetcode_Roman_To_Int
}

// Reference: https://leetcode.com/problems/roman-to-integer/
func Leetcode_Roman_To_Int() {
	fmt.Println("Input: s = 'III'")
	fmt.Println("Output:", romanToInt("III"))
	fmt.Println("Input: s = 'LVIII'")
	fmt.Println("Output:", romanToInt("LVIII"))
	fmt.Println("Input: s = 'MCMXCIV'")
	fmt.Println("Output:", romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	res, last := 0, 0
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		if cur < last {
			res -= cur
		} else {
			res += cur
		}
		last = cur
	}
	return res
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/integer-to-roman/
func init() {
	Solutions[12] = func() {
		fmt.Println("Input: num = 58")
		fmt.Println("Output:", intToRoman(58))
		fmt.Println("Input: num = 1994")
		fmt.Println("Output:", intToRoman(1994))
	}
}

func intToRoman(num int) string {
	res := []byte{}
	k, v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}, []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(k); i++ {
		if num < k[i] {
			continue
		}

		getRoman(&res, num/k[i], v[i])
		num %= k[i]
		if num == 0 {
			break
		}
	}
	return string(res)
}

func getRoman(res *[]byte, n int, b string) {
	r := []byte{}
	for i := 0; i < n; i++ {
		r = append(r, b...)
	}
	*res = append(*res, r...)
}

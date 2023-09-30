package medium

import "fmt"

// Reference: https://leetcode.com/problems/decoded-string-at-index/
func init() {
	Solutions[880] = func() {
		fmt.Println("Input: s = 'leet2code3', k = 10")
		fmt.Println("Output:", decodeAtIndex("leet2code3", 10))
		fmt.Println("Input: s = 'a2345678999999999999999', k = 1")
		fmt.Println("Output:", decodeAtIndex("a2345678999999999999999", 1))
	}
}

func decodeAtIndex(s string, k int) string {
	l := int64(0)
	i := 0
	for l < int64(k) {
		if s[i] >= '0' && s[i] <= '9' {
			l *= int64(s[i] - '0')
		} else {
			l++
		}
		i++
	}

	for j := i - 1; j >= 0; j-- {
		if s[j] >= '0' && s[j] <= '9' {
			l /= int64(s[j] - '0')
			k %= int(l)
		} else {
			if k == 0 || k == int(l) {
				return string(s[j])
			}
			l--
		}
	}
	return ""
}

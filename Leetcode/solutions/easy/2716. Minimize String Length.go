package easy

import "fmt"

// Reference: https://leetcode.com/problems/minimize-string-length/
func init() {
	Solutions[2716] = func() {
		fmt.Println("Input: s = 'aaabc'")
		fmt.Println("Output:", minimizedStringLength("aaabc"))
		fmt.Println("Input: s = 'dddaaa'")
		fmt.Println("Output:", minimizedStringLength("dddaaa"))
	}
}

func minimizedStringLength(s string) int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	return len(m)
}

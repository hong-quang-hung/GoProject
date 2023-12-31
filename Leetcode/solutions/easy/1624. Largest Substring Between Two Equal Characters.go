package easy

import "fmt"

// Reference: https://leetcode.com/problems/largest-substring-between-two-equal-characters/
func init() {
	Solutions[1624] = func() {
		fmt.Println(`Input: s = "aa"`)
		fmt.Println(`Output:`, maxLengthBetweenEqualCharacters("aa"))
		fmt.Println(`Input: s = "abca"`)
		fmt.Println(`Output:`, maxLengthBetweenEqualCharacters("abca"))
		fmt.Println(`Input: s = "cbzxy"`)
		fmt.Println(`Output:`, maxLengthBetweenEqualCharacters("cbzxy"))
	}
}

func maxLengthBetweenEqualCharacters(s string) int {
	m := make([]int, 26)
	for i := 0; i < 26; i++ {
		m[i] = -1
	}

	res := -1
	for i, ch := range s {
		key := int(ch - 'a')
		if m[key] != -1 {
			res = max(res, i-m[key]-1)
		} else {
			m[key] = i
		}
	}
	return res
}

package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/first-unique-character-in-a-string/
func init() {
	Solutions[387] = func() {
		fmt.Println(`Input: s = "leetcode"`)
		fmt.Println(`Output:`, firstUniqChar("leetcode"))
		fmt.Println(`Input: s = "loveleetcode"`)
		fmt.Println(`Output:`, firstUniqChar("loveleetcode"))
	}
}

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, ch := range s {
		m[int(ch-'a')]++
	}

	for i, ch := range s {
		if m[int(ch-'a')] == 1 {
			return i
		}
	}
	return -1
}

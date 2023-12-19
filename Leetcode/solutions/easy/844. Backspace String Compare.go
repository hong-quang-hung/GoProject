package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/backspace-string-compare/
func init() {
	Solutions[844] = func() {
		fmt.Println(`Input: s = "ab#c", t = "ad#c"`)
		fmt.Println(`Output:`, backspaceCompare(`ab#c`, `ad#c`))
		fmt.Println(`Input: s = "ab##", t = "c#d#"`)
		fmt.Println(`Output:`, backspaceCompare(`ab##`, `c#d#`))
		fmt.Println(`Input: s = "a#c", t = "b"`)
		fmt.Println(`Output:`, backspaceCompare(`a#c`, `b`))
	}
}

func backspaceCompare(s string, t string) bool {
	rs, rt := make([]rune, 0), make([]rune, 0)
	for _, ch := range s {
		if ch == '#' {
			if len(rs) > 0 {
				rs = rs[:len(rs)-1]
			}
		} else {
			rs = append(rs, ch)
		}
	}

	for _, ch := range t {
		if ch == '#' {
			if len(rt) > 0 {
				rt = rt[:len(rt)-1]
			}
		} else {
			rt = append(rt, ch)
		}
	}
	return string(rs) == string(rt)
}

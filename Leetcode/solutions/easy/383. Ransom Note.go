package easy

import "fmt"

// Reference: https://leetcode.com/problems/ransom-note/
func init() {
	Solutions[383] = func() {
		fmt.Println(`Input: ransomNote = "a", magazine = "b"`)
		fmt.Println(`Output:`, canConstruct(`a`, `b`))
		fmt.Println(`Input: ransomNote = "aa", magazine = "ab"`)
		fmt.Println(`Output:`, canConstruct(`aa`, `ab`))
	}
}

func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}
	for _, ch := range magazine {
		m[int(ch-'a')]++
	}

	for _, ch := range ransomNote {
		m[int(ch-'a')]--
		if m[int(ch-'a')] < 0 {
			return false
		}
	}
	return true
}

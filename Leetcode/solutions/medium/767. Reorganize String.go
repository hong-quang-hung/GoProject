package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/reorganize-string/
func init() {
	Solutions[767] = func() {
		fmt.Println("Input: s = 'aab'")
		fmt.Println("Output:", reorganizeString("aab"))
		fmt.Println("Input: s = 'aaab'")
		fmt.Println("Output:", reorganizeString("aaab"))
	}
}

func reorganizeString(s string) string {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	sortedChars := []rune{}
	for ch := range m {
		sortedChars = append(sortedChars, ch)
	}

	sort.Slice(sortedChars, func(i, j int) bool {
		return m[sortedChars[i]] > m[sortedChars[j]]
	})

	if m[sortedChars[0]] > (len(s)+1)/2 {
		return ""
	}

	i := 0
	res := make([]rune, len(s))
	for _, ch := range sortedChars {
		for j := 0; j < m[ch]; j++ {
			if i >= len(s) {
				i = 1
			}
			res[i] = ch
			i += 2
		}
	}
	return string(res)
}

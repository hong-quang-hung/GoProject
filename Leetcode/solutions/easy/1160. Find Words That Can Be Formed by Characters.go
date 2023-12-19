package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/
func init() {
	Solutions[1160] = func() {
		fmt.Println(`Input: words = ["cat","bt","hat","tree"], chars = "atach"`)
		fmt.Println(`Output:`, countCharacters([]string{`cat`, `bt`, `hat`, `tree`}, `atach`))
		fmt.Println(`Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"`)
		fmt.Println(`Output:`, countCharacters([]string{`hello`, `world`, `leetcode`}, `welldonehoneyr`))
	}
}

func countCharacters(words []string, chars string) int {
	m := make([]int, 26)
	for _, ch := range chars {
		m[int(ch-'a')]++
	}

	f := func(w string) bool {
		h := make([]int, 26)
		copy(h, m)
		for _, ch := range w {
			i := int(ch - 'a')
			h[i]--
			if h[i] < 0 {
				return false
			}
		}
		return true
	}

	res := 0
	for _, word := range words {
		if f(word) {
			res += len(word)
		}
	}
	return res
}

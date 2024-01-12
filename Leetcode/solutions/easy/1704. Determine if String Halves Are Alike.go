package easy

import "fmt"

// Reference: https://leetcode.com/problems/determine-if-string-halves-are-alike/
func init() {
	Solutions[1704] = func() {
		fmt.Println(`Input: s = "book"`)
		fmt.Println(`Output:`, halvesAreAlike("book"))
		fmt.Println(`Input: s = "textbook"`)
		fmt.Println(`Output:`, halvesAreAlike("textbook"))
	}
}

func halvesAreAlike(s string) bool {
	vowels := "aAeEiIoOuU"
	isVowels := func(c byte) bool {
		for i := range vowels {
			if vowels[i] == c {
				return true
			}
		}
		return false
	}

	n := len(s)
	left, right := 0, 0
	for i := 0; i*2 < n; i++ {
		if isVowels(s[i]) {
			left++
		}
		if isVowels(s[n-1-i]) {
			right++
		}
	}
	return right == left
}

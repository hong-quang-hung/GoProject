package easy

import "fmt"

// Reference: https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/
func init() {
	Solutions[2586] = func() {
		fmt.Println(`Input: words = ["are","amy","u"], left = 0, right = 2`)
		fmt.Println(`Output:`, vowelStrings([]string{`hey`, `aeo`, `mu`, `ooo`, `artro`}, 1, 4))
	}
}

func vowelStrings(words []string, left int, right int) int {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'u' || c == 'o'
	}
	var res int = 0
	for i := left; i <= right; i++ {
		w := words[i]
		l := len(w)
		if (l == 1 && isVowel(w[0])) || (isVowel(w[0]) && isVowel(w[l-1])) {
			res++
		}
	}
	return res
}

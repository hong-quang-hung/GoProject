package easy

import "fmt"

func init() {
	Solutions[2586] = Leetcode_Vowel_Strings
}

// Reference: https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/
func Leetcode_Vowel_Strings() {
	fmt.Println("Input: words = ['hey', 'aeo', 'mu', 'ooo', 'artro'], left = 1, right = 4")
	fmt.Println("Output:", vowelStrings([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4))
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

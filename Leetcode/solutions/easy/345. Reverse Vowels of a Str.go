package easy

import "fmt"

// Reference: https://leetcode.com/problems/reverse-vowels-of-a-string/
func init() {
	Solutions[345] = func() {
		fmt.Println("Input: s = 'hello'")
		fmt.Println("Output:", reverseVowels("hello"))
		fmt.Println("Input: s = 'leetcode'")
		fmt.Println("Output:", reverseVowels("leetcode"))
		fmt.Println("Input: s = 'aA'")
		fmt.Println("Output:", reverseVowels("aA"))
	}
}

func reverseVowels(s string) string {
	res := []byte(s)
	m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !m[s[i]] {
			i++
		}

		for i < j && !m[s[j]] {
			j--
		}

		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
func Leetcode_Max_Vowels() {
	fmt.Println("Input: s = 'abciiidef', k = 3")
	fmt.Println("Output:", maxVowels("abciiidef", 3))
	fmt.Println("Input: s = 'aeiou', k = 2")
	fmt.Println("Output:", maxVowels("aeiou", 2))
	fmt.Println("Input: s = 'leetcode', k = 3")
	fmt.Println("Output:", maxVowels("leetcode", 3))
}

func maxVowels(s string, k int) int {
	res := 0
	m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for i := 0; i < k; i++ {
		if m[s[i]] {
			res++
		}
	}

	count := res
	for i := k; i < len(s); i++ {
		if m[s[i-k]] == m[s[i]] {
			continue
		} else if m[s[i]] && !m[s[i-k]] {
			count++
		} else {
			count--
		}
		res = max(res, count)
	}
	return res
}

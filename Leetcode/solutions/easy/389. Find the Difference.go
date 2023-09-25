package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-the-difference/
func init() {
	Solutions[389] = func() {
		fmt.Println("Input: s = 'abcd', t = 'abcde'")
		fmt.Println("Output:", findTheDifference("abcd", "abcde"))
		fmt.Println("Input: s = '', t = 'y'")
		fmt.Println("Output:", findTheDifference("", "y"))
	}
}

func findTheDifference(s string, t string) byte {
	m := [26]int{}
	for i := range s {
		m[int(s[i]-'a')]++
	}

	for i, ch := range t {
		idx := int(ch - 'a')
		if m[idx] == 0 {
			return t[i]
		}
		m[idx]--
	}
	return ' '
}

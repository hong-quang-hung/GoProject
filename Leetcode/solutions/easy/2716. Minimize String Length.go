package easy

import "fmt"

func init() {
	Solutions[2716] = Leetcode_Minimized_String_Length
}

// Reference: https://leetcode.com/problems/minimize-string-length/
func Leetcode_Minimized_String_Length() {
	fmt.Println("Input: s = 'aaabc'")
	fmt.Println("Output:", minimizedStringLength("aaabc"))
	fmt.Println("Input: s = 'cbbd'")
	fmt.Println("Output:", minimizedStringLength("cbbd"))
	fmt.Println("Input: s = 'dddaaa'")
	fmt.Println("Output:", minimizedStringLength("dddaaa"))
}

func minimizedStringLength(s string) int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	return len(m)
}

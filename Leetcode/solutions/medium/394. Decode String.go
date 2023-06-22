package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/decode-string/
func Leetcode_Decode_String() {
	fmt.Println("Input: s = '3[a]2[bc]'")
	fmt.Println("Output:", decodeString("3[a]2[bc]"))
	fmt.Println("Input: s = '3[a2[c]]'")
	fmt.Println("Output:", decodeString("3[a2[c]]"))
	fmt.Println("Input: s = '2[abc]3[cd]ef'")
	fmt.Println("Output:", decodeString("2[abc]3[cd]ef"))
}

func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}

	if s[0]-'0' <= 9 {
		i, n := 1, int(s[0]-'0')
		for i < len(s) && s[i] != '[' {
			n = n*10 + int(s[i]-'0')
		}

		j := i + 1
		stack := []byte{s[i]}
		for len(stack) != 0 {
			if s[j] == '[' {
				stack = append(stack, s[j])
			} else if s[j] == ']' {
				stack = stack[:len(stack)-1]
			}
			j++
		}

		res := ""
		left := decodeString(s[i+1 : j-1])
		for r := 0; r < n; r++ {
			res += left
		}
		return res + decodeString(s[j:])
	}

	i := 0
	for i < len(s) && s[i]-'0' > 9 {
		i++
	}
	return s[0:i] + decodeString(s[i:])
}

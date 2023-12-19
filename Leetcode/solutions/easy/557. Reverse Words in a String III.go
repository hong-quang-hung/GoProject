package easy

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/reverse-words-in-a-string-iii/
func init() {
	Solutions[557] = func() {
		fmt.Println(`Input: s = "Let's take LeetCode contest"`)
		fmt.Println(`Output:`, reverseWords(`Let's take LeetCode contest`))
		fmt.Println(`Input: s = "Mr Ding"`)
		fmt.Println(`Output:`, reverseWords(`God Ding`))
	}
}

func reverseWords(s string) string {
	words := strings.Split(s, ` `)
	res := []string{}
	for _, word := range words {
		temp := []byte(word)
		for i, j := 0, len(temp)-1; i < j; {
			temp[i], temp[j] = temp[j], temp[i]
			i++
			j--
		}
		res = append(res, string(temp))
	}
	return strings.Join(res, ` `)
}

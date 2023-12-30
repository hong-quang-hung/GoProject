package easy

import "fmt"

// Reference: https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
func init() {
	Solutions[1897] = func() {
		fmt.Println(`Input: words = ["abc","aabc","bc"]`)
		fmt.Println(`Output:`, makeEqual([]string{"abc", "aabc", "bc"}))
		fmt.Println(`Input: words = ["ab","a"]`)
		fmt.Println(`Output:`, makeEqual([]string{"ab", "a"}))
	}
}

func makeEqual(words []string) bool {
	n := len(words)
	m := make(map[int]int)
	for _, word := range words {
		for _, ch := range word {
			m[int(ch-'a')]++
		}
	}

	for _, val := range m {
		if val%n != 0 {
			return false
		}
	}
	return true
}

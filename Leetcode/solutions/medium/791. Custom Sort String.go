package medium

import "fmt"

// Reference: https://leetcode.com/problems/custom-sort-string/
func init() {
	Solutions[791] = func() {
		fmt.Println(`Input: order = "cba", s = "abcd"`)
		fmt.Println(`Output:`, customSortString("cba", "abcd"))
		fmt.Println(`Input: order = "bcafg", s = "abcd"`)
		fmt.Println(`Output:`, customSortString("bcafg", "abcd"))
	}
}

func customSortString(order string, s string) string {
	m := make([]int, 26)
	for _, ch := range s {
		m[int(ch-'a')]++
	}

	res := make([]rune, 0)
	for _, ch := range order {
		k := int(ch - 'a')
		for i := 0; i < m[k]; i++ {
			res = append(res, ch)
		}
		m[k] = 0
	}

	for i := 0; i < 26; i++ {
		if m[i] > 0 {
			for j := 0; j < m[i]; j++ {
				res = append(res, rune(i+'a'))
			}
		}
	}
	return string(res)
}

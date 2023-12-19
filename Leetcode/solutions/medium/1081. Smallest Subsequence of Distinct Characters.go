package medium

import "fmt"

// Reference: https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
func init() {
	Solutions[1081] = func() {
		fmt.Println(`Input: s = "bcabc"`)
		fmt.Println(`Output:`, smallestSubsequence(`bcabc`))
		fmt.Println(`Input: s = "cbacdcbc"`)
		fmt.Println(`Output:`, smallestSubsequence(`cbacdcbc`))
	}
}

func smallestSubsequence(s string) string {
	left := ['z' + 1]int{}
	for _, c := range s {
		left[c]++
	}

	res := []rune{}
	m := map[rune]struct{}{}
	for _, c := range s {
		left[c]--
		if _, ok := m[c]; ok {
			continue
		}

		for len(res) > 0 && c < res[len(res)-1] && left[res[len(res)-1]] > 0 {
			x := res[len(res)-1]
			res = res[:len(res)-1]
			delete(m, x)
		}

		res = append(res, c)
		m[c] = struct{}{}
	}
	return string(res)
}

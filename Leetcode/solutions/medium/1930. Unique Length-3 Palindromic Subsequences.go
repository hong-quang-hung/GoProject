package medium

import "fmt"

// Reference: https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
func init() {
	Solutions[1930] = func() {
		fmt.Println(`Input: s = "aabca"`)
		fmt.Println(`Output:`, countPalindromicSubsequence(`aabca`))
		fmt.Println(`Input: s = "adc"`)
		fmt.Println(`Output:`, countPalindromicSubsequence(`adc`))
		fmt.Println(`Input: s = "bbcbaba"`)
		fmt.Println(`Output:`, countPalindromicSubsequence(`bbcbaba`))
	}
}

func countPalindromicSubsequence(str string) int {
	pos := [26][2]int{}
	for i, s := range str {
		offset := s - 97
		if pos[offset][0] == 0 {
			pos[offset][0] = i
		}
		if pos[offset][1] == 0 {
			pos[offset][1] = i
		}
		if pos[offset][0] > i {
			pos[offset][0] = i
		}
		if pos[offset][1] < i {
			pos[offset][1] = i
		}
	}
	pos[str[0]-97][0] = 0

	res := 0
	for _, p := range pos {
		if p[0] == p[1] {
			continue
		}
		counter := [26]int{}
		for i := p[0] + 1; i < p[1]; i++ {
			counter[str[i]-97] = 1
		}
		for _, c := range counter {
			res += c
		}
	}
	return res
}

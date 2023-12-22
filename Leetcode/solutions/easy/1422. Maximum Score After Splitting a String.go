package easy

import "fmt"

// Reference: https://leetcode.com/problems/maximum-score-after-splitting-a-string/
func init() {
	Solutions[1422] = func() {
		fmt.Println(`Input: s = "011101"`)
		fmt.Println(`Output:`, maxScore("011101"))
		fmt.Println(`Input: s = "00111"`)
		fmt.Println(`Output:`, maxScore("00111"))
		fmt.Println(`Input: s = "1111"`)
		fmt.Println(`Output:`, maxScore("1111"))
	}
}

func maxScore(s string) int {
	one := 0
	for _, ch := range s {
		if ch == '1' {
			one++
		}
	}

	res := 0
	zero := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zero++
		} else {
			one--
		}
		res = max(res, zero+max(one, 0))
	}
	return res
}

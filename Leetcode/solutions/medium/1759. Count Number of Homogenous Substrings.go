package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-number-of-homogenous-substrings/
func init() {
	Solutions[1759] = func() {
		fmt.Println(`Input: s = "abbcccaa"`)
		fmt.Println(`Output:`, countHomogenous(`abbcccaa`))
		fmt.Println(`Input: s = "xy"`)
		fmt.Println(`Output:`, countHomogenous(`xy`))
		fmt.Println(`Input: s = "zzzzz"`)
		fmt.Println(`Output:`, countHomogenous(`zzzzz`))
	}
}

func countHomogenous(s string) int {
	res, i, j := 0, 0, 0
	for i <= j && j < len(s) {
		for j < len(s) && s[i] == s[j] {
			j++
		}

		temp := j - i
		res += ((temp%mod + 1) * temp / 2) % mod
		i = j
	}
	return res
}

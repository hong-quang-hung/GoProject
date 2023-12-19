package medium

import "fmt"

// Reference: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func init() {
	Solutions[17] = func() {
		fmt.Println(`Input: digits = '23'`)
		fmt.Println(`Output:`, letterCombinations(`23`))
		fmt.Println(`Input: digits = ''`)
		fmt.Println(`Output:`, letterCombinations(``))
		fmt.Println(`Input: digits = '2'`)
		fmt.Println(`Output:`, letterCombinations(`2`))
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := map[byte]string{'2': `abc`, '3': `def`, '4': `ghi`, '5': `jkl`, '6': `mno`, '7': `pqrs`, '8': `tuv`, '9': `wxyz`}
	res := make([]string, 0)
	letterCombinationsSolved(digits, m, &res, 0, ``)
	return res
}

func letterCombinationsSolved(digits string, m map[byte]string, res *[]string, i int, s string) {
	if i >= len(digits) {
		*res = append(*res, s)
		return
	}

	for _, ch := range m[digits[i]] {
		letterCombinationsSolved(digits, m, res, i+1, s+string(ch))
	}
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-all-anagrams-in-a-string/
func init() {
	Solutions[438] = func() {
		fmt.Println("Input: s = 'cbaebabacd', p = 'abc'")
		fmt.Println("Output:", findAnagrams("cbaebabacd", "abc"))
		fmt.Println("Input: s = 'abab', p = 'abc'")
		fmt.Println("Output:", findAnagrams("abab", "ab"))
	}
}

func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return nil
	}

	mS, mP, res := make([]int, 26), make([]int, 26), []int{}
	for i := 0; i < pLen; i++ {
		mP[p[i]-'a']++
		mS[s[i]-'a']++
	}

	if compareSlice(mS, mP) {
		res = append(res, 0)
	}

	for i := pLen; i < sLen; i++ {
		mS[s[i-pLen]-'a']--
		mS[s[i]-'a']++

		if compareSlice(mS, mP) {
			res = append(res, i-pLen+1)
		}
	}
	return res
}

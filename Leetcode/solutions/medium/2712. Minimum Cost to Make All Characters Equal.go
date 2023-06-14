package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-cost-to-make-all-characters-equal/
func Leetcode_Minimum_Cost_II() {
	fmt.Println("Input: num = '0011'")
	fmt.Println("Output:", minimumCost_ii("0011"))
	fmt.Println("Input: num = '010101'")
	fmt.Println("Output:", minimumCost_ii("010101"))
}

func minimumCost_ii(s string) int64 {
	res, n := int64(0), len(s)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			res += int64(min(i, n-i))
		}
	}
	return res
}

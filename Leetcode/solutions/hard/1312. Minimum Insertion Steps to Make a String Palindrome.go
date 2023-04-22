package hard

import "fmt"

// Reference: https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
func Leetcode_Min_Insertions() {
	fmt.Println("Input: s = 'zzazz'")
	fmt.Println("Output:", minInsertions("zzazz"))
	fmt.Println("Input: s = 'mbadm'")
	fmt.Println("Output:", minInsertions("mbadm"))
	fmt.Println("Input: s = 'leetcode'")
	fmt.Println("Output:", minInsertions("leetcode"))
}

func minInsertions(s string) int {
	n := len(s)
	dp, dpPrev := make([]int, n), make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[j] = dpPrev[j-1] + 2
			} else {
				dp[j] = max(dpPrev[j], dp[j-1])
			}
		}
		copy(dpPrev, dp)
	}
	return n - dp[n-1]
}

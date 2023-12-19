package hard

import "fmt"

// Reference: https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
func init() {
	Solutions[1639] = func() {
		fmt.Println(`Input: words = ["acca","bbbb","caca"], target = "aba"`)
		fmt.Println(`Output:`, numWaysII([]string{`acca`, `bbbb`, `caca`}, `aba`))
		fmt.Println(`Input: words = ["abba","baab"], target = "bab"`)
		fmt.Println(`Output:`, numWaysII([]string{`abba`, `baab`}, `bab`))
	}
}

func numWaysII(words []string, target string) int {
	n, m := len(target), len(words[0])
	dp := make([][]int64, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int64, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = -1
		}
	}

	cnt := make([][]int, 26)
	for i := 0; i < 26; i++ {
		cnt[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		for _, word := range words {
			cnt[word[j]-'a'][j]++
		}
	}
	return int(numWaysIISolved(dp, cnt, target, n, m))
}

func numWaysIISolved(dp [][]int64, cnt [][]int, target string, i int, j int) int64 {
	if j == 0 {
		if i == 0 {
			return 1
		} else {
			return 0
		}
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	dp[i][j] = numWaysIISolved(dp, cnt, target, i, j-1)
	if i > 0 {
		dp[i][j] += int64(cnt[target[i-1]-'a'][j-1]) * int64(numWaysIISolved(dp, cnt, target, i-1, j-1))
	}

	dp[i][j] %= 1000000007
	return dp[i][j]
}

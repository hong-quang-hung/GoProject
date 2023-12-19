package hard

import "fmt"

// Reference: https://leetcode.com/problems/shortest-common-supersequence/
func init() {
	Solutions[1092] = func() {
		fmt.Println(`Input: str1 = "abac", str2 = "cab"`)
		fmt.Println(`Output:`, shortestCommonSupersequence(`abac`, `cab`))
		fmt.Println(`Input: str1 = "aaaaaaaa", str2 = "aaaaaaaa"`)
		fmt.Println(`Output:`, shortestCommonSupersequence(`aaaaaaaa`, `aaaaaaaa`))
		fmt.Println(`Input: str1 = "cddb", str2 = "cab"`)
		fmt.Println(`Output:`, shortestCommonSupersequence(`cddb`, `cab`))
	}
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	if m < n {
		str1, str2 = str2, str1
		m, n = n, m
	}

	dp := make([][]string, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]string, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(str1[i-1])
			} else {
				if len(dp[i-1][j]) > len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	a, b, lcs, res := 0, 0, dp[m][n], []byte{}
	for i := range lcs {
		for str1[a] != lcs[i] {
			res = append(res, str1[a])
			a++
		}

		for str2[b] != lcs[i] {
			res = append(res, str2[b])
			b++
		}

		res = append(res, lcs[i])
		a++
		b++
	}
	return string(res) + str1[a:] + str2[b:]
}

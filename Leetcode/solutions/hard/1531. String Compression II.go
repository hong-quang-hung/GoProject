package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/string-compression-ii/
func init() {
	Solutions[1531] = func() {
		fmt.Println(`Input: s = "aaabcccd", k = 2`)
		fmt.Println(`Output:`, getLengthOfOptimalCompression("aaabcccd", 2))
		fmt.Println(`Input: s = "aabbaa", k = 2`)
		fmt.Println(`Output:`, getLengthOfOptimalCompression("aabbaa", 2))
		fmt.Println(`Input: s = "aaaaaaaaaaa", k = 0`)
		fmt.Println(`Output:`, getLengthOfOptimalCompression("aaaaaaaaaaa", 0))
	}
}

func getLengthOfOptimalCompression(s string, k int) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	l := func(i int) int {
		if i == 1 {
			return 1
		} else if i <= 9 {
			return 2
		} else if i <= 99 {
			return 3
		}
		return 4
	}

	var f func(str string, i int) int
	f = func(newS string, newK int) int {
		n := len(newS)
		if dp[n][newK] != -1 {
			return dp[n][newK]
		}

		if n <= newK {
			return 0
		}

		res := math.MaxInt32
		if newK > 0 {
			res = f(newS[1:], newK-1)
		}

		count := 0
		currentK := newK
		for i := range newS {
			if newS[i] == newS[0] {
				count++
			} else if currentK > 0 {
				currentK--
			} else {
				break
			}
			res = min(res, l(count)+f(newS[i+1:], currentK))
		}

		dp[n][newK] = res
		return res
	}
	return f(s, k)
}

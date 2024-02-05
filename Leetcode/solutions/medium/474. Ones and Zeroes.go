package medium

import "fmt"

// Reference: https://leetcode.com/problems/ones-and-zeroes/
func init() {
	Solutions[474] = func() {
		fmt.Println(`Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3`)
		fmt.Println(`Output:`, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
		fmt.Println(`Input: strs = ["10","0","1"], m = 1, n = 1`)
		fmt.Println(`Output:`, findMaxForm([]string{"10", "0", "1"}, 1, 1))
		fmt.Println(`Input: strs = ["10","0001","111001","1","0"], m = 4, n = 3`)
		fmt.Println(`Output:`, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 4, 3))
	}
}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	zero := make(map[int]int, l)
	one := make(map[int]int, l)
	for i, s := range strs {
		ch0, ch1 := 0, 0
		for _, ch := range s {
			if ch == '0' {
				ch0++
			} else {
				ch1++
			}
		}
		zero[i], one[i] = ch0, ch1
	}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := range strs {
		zeros, ones := zero[i], one[i]
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zeros][j-ones])
			}
		}
	}
	return dp[m][n]
}

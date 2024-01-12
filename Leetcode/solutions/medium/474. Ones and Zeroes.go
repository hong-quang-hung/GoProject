package medium

import "fmt"

// Reference: https://leetcode.com/problems/ones-and-zeroes/
func init() {
	Solutions[474] = func() {
		fmt.Println(`Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3`)
		fmt.Println(`Output:`, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
		fmt.Println(`Input: strs = ["10","0","1"], m = 1, n = 1`)
		fmt.Println(`Output:`, findMaxForm([]string{"10", "0", "1"}, 1, 1))
	}
}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
			for z := 0; z <= n; z++ {
				dp[i][j][z] = -1
			}
		}
	}

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

	var f func(i, z, o int) int
	f = func(i, z, o int) int {
		if z < 0 || o < 0 || i >= l {
			return 0
		}

		if z == 0 && o == 0 {
			return 1
		}

		if dp[i][z][o] != -1 {
			return dp[i][z][o]
		}

		dp[i][z][o] = max(f(i+1, z, o), f(i+1, z-zero[i], o-one[i])+1)
		return dp[i][z][o]
	}
	return f(0, m, n)
}

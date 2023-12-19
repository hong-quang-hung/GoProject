package hard

import "fmt"

// Reference: https://leetcode.com/problems/profitable-schemes/
func init() {
	Solutions[879] = func() {
		fmt.Println(`Input: n = 5, minProfit = 3, group = [2,2], profit = [2,3]`)
		fmt.Println(`Output:`, profitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
		fmt.Println(`Input: n = 10, minProfit = 5, group = [2,3,5], profit = [6,7,8]`)
		fmt.Println(`Output:`, profitableSchemes(10, 5, []int{2, 3, 5}, []int{6, 7, 8}))
	}
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	m := len(group)
	dp := make([][][]int, 101)
	for i := 0; i < 101; i++ {
		dp[i] = make([][]int, 101)
		for j := 0; j < 101; j++ {
			dp[i][j] = make([]int, 101)
		}
	}

	for i := 0; i <= n; i++ {
		dp[m][i][minProfit] = 1
	}

	for index := m - 1; index >= 0; index-- {
		for count := 0; count <= n; count++ {
			for pro := 0; pro <= minProfit; pro++ {
				dp[index][count][pro] = dp[index+1][count][pro]
				if count+group[index] <= n {
					dp[index][count][pro] = (dp[index][count][pro] + dp[index+1][count+group[index]][min(minProfit, pro+profit[index])]) % mod
				}
			}
		}
	}
	return dp[0][0][0]
}

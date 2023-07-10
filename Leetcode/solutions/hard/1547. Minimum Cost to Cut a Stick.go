package hard

import (
	"fmt"
	"math"
	"sort"
)

func init() {
	Solutions[1547] = Leetcode_Min_Cost_Cut
}

// Reference: https://leetcode.com/problems/minimum-cost-to-cut-a-stick/
func Leetcode_Min_Cost_Cut() {
	fmt.Println("Input: n = 7, cuts = [1,3,4,5]")
	fmt.Println("Output:", minCostCut(7, []int{1, 3, 4, 5}))
	fmt.Println("Input: n = 9, cuts = [5,6,1,4,2]")
	fmt.Println("Output:", minCostCut(9, []int{5, 6, 1, 4, 2}))
}

func minCostCut(n int, cuts []int) int {
	cuts = append(cuts, []int{0, n}...)
	sort.Ints(cuts)

	l := len(cuts)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		for j := 0; j < l; j++ {
			dp[i][j] = -1
		}
	}
	return minCostCutSolved(cuts, dp, 0, l-1)
}

func minCostCutSolved(cuts []int, dp [][]int, i, j int) int {
	if i-j == 1 || j-i == 1 || i == j {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	m := math.MaxInt32
	for k := i + 1; k < j; k++ {
		m = min(m, minCostCutSolved(cuts, dp, i, k)+minCostCutSolved(cuts, dp, k, j)+cuts[j]-cuts[i])
	}

	dp[i][j] = m
	return m
}

package hard

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/minimize-the-total-price-of-the-trips/
func Leetcode_Minimum_Total_Price() {
	fmt.Println("Input: n = 4, edges = [[0,1],[1,2],[1,3]], price = [2,2,10,6], trips = [[0,3],[2,1],[2,3]]")
	fmt.Println("Output:", minimumTotalPrice(4, utils.S2SoSliceInt("[[0,1],[1,2],[1,3]]"), []int{2, 2, 10, 6}, utils.S2SoSliceInt("[[0,3],[2,1],[2,3]]")))
	fmt.Println("Input: n = 2, edges = [[0,1]], price = [2,2], trips = [[0,0]]")
	fmt.Println("Output:", minimumTotalPrice(2, utils.S2SoSliceInt("[[0,1]]"), []int{2, 2}, utils.S2SoSliceInt("[[0,0]]")))
	fmt.Println("Input: n = 5, edges = [[1,2],[2,0],[0,3],[3,4]], price = [12,26,22,12,2], trips = [[3,3],[3,2],[3,0],[3,4],[1,1]]")
	fmt.Println("Output:", minimumTotalPrice(5, utils.S2SoSliceInt("[[1,2],[2,0],[0,3],[3,4]]"), []int{12, 26, 22, 12, 2}, utils.S2SoSliceInt("[[3,3],[3,2],[3,0],[3,4],[1,1]]")))
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	dp := make([][][2]int, n)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}

	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dp[i][j][0] = price[i]
				dp[i][j][1] = price[i] / 2
			} else {
				dp[i][j][0] = -1
				dp[i][j][1] = -1
			}
		}
	}

	res1, res2 := 0, 0
	for _, trip := range trips {
		res1 += minimumTotalPriceDFS(g, price, dp, make([]bool, n), trip[0], trip[1], 0)
		res2 += minimumTotalPriceDFS(g, price, dp, make([]bool, n), trip[0], trip[1], 1)
	}
	return min(res1, res2)
}

func minimumTotalPriceDFS(g [][]int, price []int, dp [][][2]int, visited []bool, start int, end int, half int) int {
	if visited[start] {
		return 500001
	}

	if dp[start][end][half] != -1 || dp[end][start][half] != -1 {
		return dp[start][end][half]
	}

	minP := 500001
	visited[start] = true
	for _, v := range g[start] {
		if half == 0 {
			minP = min(minP, price[start]+min(minimumTotalPriceDFS(g, price, dp, visited, v, end, 1), minimumTotalPriceDFS(g, price, dp, visited, v, end, 0)))
		} else {
			minP = min(minP, minimumTotalPriceDFS(g, price, dp, visited, v, end, 0)+price[start]/2)
		}
	}

	if minP < 500001 {
		dp[start][end][half] = minP
		dp[end][start][half] = minP
	}
	return minP
}

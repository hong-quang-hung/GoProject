package hard

import "fmt"

func init() {
	Solutions[2646] = Leetcode_Minimum_Total_Price
}

// Reference: https://leetcode.com/problems/minimize-the-total-price-of-the-trips/
func Leetcode_Minimum_Total_Price() {
	fmt.Println("Input: n = 4, edges = [[0,1],[1,2],[1,3]], price = [2,2,10,6], trips = [[0,3],[2,1],[2,3]]")
	fmt.Println("Output:", minimumTotalPrice(4, S2SoSliceInt("[[0,1],[1,2],[1,3]]"), []int{2, 2, 10, 6}, S2SoSliceInt("[[0,3],[2,1],[2,3]]")))
	fmt.Println("Input: n = 2, edges = [[0,1]], price = [2,2], trips = [[0,0]]")
	fmt.Println("Output:", minimumTotalPrice(2, S2SoSliceInt("[[0,1]]"), []int{2, 2}, S2SoSliceInt("[[0,0]]")))
	fmt.Println("Input: n = 5, edges = [[1,2],[2,0],[0,3],[3,4]], price = [12,26,22,12,2], trips = [[3,3],[3,2],[3,0],[3,4],[1,1]]")
	fmt.Println("Output:", minimumTotalPrice(5, S2SoSliceInt("[[1,2],[2,0],[0,3],[3,4]]"), []int{12, 26, 22, 12, 2}, S2SoSliceInt("[[3,3],[3,2],[3,0],[3,4],[1,1]]")))
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	dp := make([][2]int, n)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}

	for i := 0; i < n; i++ {
		dp[i] = [2]int{-1, -1}
	}

	cnt := make([]int, n)
	for _, trip := range trips {
		cnt[trip[0]]++
		minimumTotalPriceCountNode(g, trip[0], -1, trip[1], cnt)
	}
	return minimumTotalPriceDFS(g, price, dp, 0, -1, 0, cnt)
}

func minimumTotalPriceDFS(g [][]int, price []int, dp [][2]int, node int, par int, parTaken int, cnt []int) int {
	if dp[node][parTaken] != -1 {
		return dp[node][parTaken]
	}

	childs := 0
	for _, i := range g[node] {
		if i != par {
			childs++
		}
	}

	if childs == 0 {
		if parTaken == 1 {
			dp[node][parTaken] = cnt[node] * (price[node])
			return dp[node][parTaken]
		}

		dp[node][parTaken] = cnt[node] * (price[node] / 2)
		return dp[node][parTaken]
	}

	if parTaken == 1 {
		ans := 0
		for _, i := range g[node] {
			if i != par {
				ans += minimumTotalPriceDFS(g, price, dp, i, node, 0, cnt)
			}
		}

		ans += cnt[node] * price[node]
		dp[node][parTaken] = ans
		return dp[node][parTaken]
	}

	ans1, ans2 := 0, 0
	for _, i := range g[node] {
		if i != par {
			ans1 += minimumTotalPriceDFS(g, price, dp, i, node, 0, cnt)
			ans2 += minimumTotalPriceDFS(g, price, dp, i, node, 1, cnt)
		}
	}

	ans1 += cnt[node] * (price[node])
	ans2 += cnt[node] * (price[node] / 2)
	dp[node][parTaken] = min(ans1, ans2)
	return dp[node][parTaken]
}

func minimumTotalPriceCountNode(g [][]int, node int, par int, dst int, cnt []int) bool {
	if node == dst {
		return true
	}

	for _, i := range g[node] {
		if i != par {
			if minimumTotalPriceCountNode(g, i, node, dst, cnt) {
				cnt[i]++
				return true
			}
		}
	}
	return false
}

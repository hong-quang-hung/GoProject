package hard

import "fmt"

// Reference: https://leetcode.com/problems/count-all-possible-routes/
func init() {
	Solutions[1575] = func() {
		fmt.Println("Input: locations = [2,3,6,8,4], start = 1, finish = 3, fuel = 5")
		fmt.Println("Output:", countRoutes([]int{2, 3, 6, 8, 4}, 1, 3, 5))
		fmt.Println("Input: locations = [4,3,1], start = 1, finish = 0, fuel = 6")
		fmt.Println("Output:", countRoutes([]int{4, 3, 1}, 1, 0, 6))
		fmt.Println("Input: locations = [5,2,1], start = 0, finish = 2, fuel = 3")
		fmt.Println("Output:", countRoutes([]int{5, 2, 1}, 0, 2, 3))
	}
}

func countRoutes(locations []int, start int, finish int, fuel int) int {
	n := len(locations)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, fuel+1)
		for j := 0; j <= fuel; j++ {
			dp[i][j] = -1
		}
	}
	return countRoutesSolved(locations, start, finish, fuel, dp)
}

func countRoutesSolved(locations []int, start int, finish int, fuel int, dp [][]int) int {
	if fuel < 0 {
		return 0
	}

	if dp[start][fuel] != -1 {
		return dp[start][fuel]
	}

	count := 0
	if start == finish {
		count = 1
	}

	for idx, location := range locations {
		if idx != start {
			count = (count + countRoutesSolved(locations, idx, finish, fuel-abs(locations[start]-location), dp)) % mod
		}
	}

	dp[start][fuel] = count
	return count
}

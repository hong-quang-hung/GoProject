package hard

import "fmt"

func init() {
	Solutions[2218] = Leetcode_Max_Value_Of_Coins
}

// Reference: https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/
func Leetcode_Max_Value_Of_Coins() {
	fmt.Println("Input: piles = [[1,100,3],[7,8,9]], k = 2")
	fmt.Println("Output:", maxValueOfCoins(S2SoSliceInt("[[1,100,3],[7,8,9]]"), 2))
	fmt.Println("Input: piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7")
	fmt.Println("Output:", maxValueOfCoins(S2SoSliceInt("[[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]]"), 7))
}

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	dp := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for coins := 0; coins <= k; coins++ {
			dp[i][coins] = -1
		}
	}
	return maxValueOfCoinsFunc(piles, dp, n, k)
}

func maxValueOfCoinsFunc(piles [][]int, dp [][]int, i, coins int) int {
	if i == 0 {
		return 0
	}

	if dp[i][coins] != -1 {
		return dp[i][coins]
	}

	currentSum := 0
	for currentCoins := 0; currentCoins <= min(len(piles[i-1]), coins); currentCoins++ {
		if currentCoins > 0 {
			currentSum += piles[i-1][currentCoins-1]
		}
		dp[i][coins] = max(dp[i][coins], maxValueOfCoinsFunc(piles, dp, i-1, coins-currentCoins)+currentSum)
	}
	return dp[i][coins]
}

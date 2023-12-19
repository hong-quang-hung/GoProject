package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference/
func init() {
	Solutions[1218] = func() {
		fmt.Println(`Input: arr = [1,2,3,4], difference = 1`)
		fmt.Println(`Output:`, longestSubsequence([]int{1, 2, 3, 4}, 1))
		fmt.Println(`Input: arr = [1,3,5,7], difference = 1`)
		fmt.Println(`Output:`, longestSubsequence([]int{1, 3, 5, 7}, 1))
		fmt.Println(`Input: arr = [1,5,7,8,5,3,4,2,1], difference = -2`)
		fmt.Println(`Output:`, longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, 2))
	}
}

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	res := 0
	for _, num := range arr {
		prev := num - difference
		dp[num] = dp[prev] + 1
		if dp[num] > res {
			res = dp[num]
		}
	}
	return res
}

package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-cost-for-tickets/
func Leetcode_Min_Cost_Tickets() {
	fmt.Println("Input: days = [1,4,6,7,8,20], costs = [2,7,15]")
	fmt.Println("Output:", mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println("Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]")
	fmt.Println("Output:", mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, len(days))
	dur := []int{1, 7, 30}
	return mincostTicketsDP(days, costs, dur, dp, 0)
}

func mincostTicketsDP(days []int, costs []int, dur []int, dp []int, i int) int {
	if i >= len(dp) {
		return 0
	}

	if dp[i] != 0 {
		return dp[i]
	}

	var j int = i
	var mincost int = math.MaxInt16
	for c, cost := range costs {
		for (j < len(days)) && (days[j] < days[i]+dur[c]) {
			j++
		}
		dp := mincostTicketsDP(days, costs, dur, dp, j) + cost
		mincost = min(mincost, dp)
	}
	dp[i] = mincost
	return mincost
}

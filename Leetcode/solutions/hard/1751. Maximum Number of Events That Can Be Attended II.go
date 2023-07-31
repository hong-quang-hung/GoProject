package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/
func init() {
	Solutions[1751] = func() {
		fmt.Println("Input: events = [[1,2,4],[3,4,3],[2,3,1]], k = 2")
		fmt.Println("Output:", maxValue(S2SoSliceInt("[[1,2,4],[3,4,3],[2,3,1]]"), 2))
		fmt.Println("Input: events = [[1,2,4],[3,4,3],[2,3,10]], k = 2")
		fmt.Println("Output:", maxValue(S2SoSliceInt("[[1,2,4],[3,4,3],[2,3,10]]"), 2))
		fmt.Println("Input: events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3")
		fmt.Println("Output:", maxValue(S2SoSliceInt("[[1,1,1],[2,2,2],[3,3,3],[4,4,4]]"), 3))
	}
}

func maxValue(events [][]int, k int) int {
	n := len(events)
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}

	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })

	for curr := n - 1; curr >= 0; curr-- {
		next := sort.Search(n, func(i int) bool { return events[i][0] > events[curr][1] })
		for count := 1; count <= k; count++ {
			dp[count][curr] = max(dp[count][curr+1], events[curr][2]+dp[count-1][next])
		}
	}
	return dp[k][0]
}

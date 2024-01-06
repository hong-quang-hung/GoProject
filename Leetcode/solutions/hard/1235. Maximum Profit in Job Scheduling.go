package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-profit-in-job-scheduling/
func init() {
	Solutions[1235] = func() {
		fmt.Println(`Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]`)
		fmt.Println(`Output:`, jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
		fmt.Println(`Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]`)
		fmt.Println(`Output:`, jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
		fmt.Println(`Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]`)
		fmt.Println(`Output:`, jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}))
	}
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	dp := make([]int, n+1)
	job := make([][3]int, n)
	for i := 0; i < n; i++ {
		job[i][0], job[i][1], job[i][2] = startTime[i], endTime[i], profit[i]
	}

	sort.Slice(job, func(i, j int) bool { return job[i][0] < job[j][0] })

	for i := n - 1; i >= 0; i-- {
		dp[i] = job[i][2]
		index := sort.Search(n, func(j int) bool { return job[j][0] >= job[i][1] })
		if index < n {
			dp[i] += dp[index]
		}
		dp[i] = max(dp[i], dp[i+1])
	}
	return dp[0]
}

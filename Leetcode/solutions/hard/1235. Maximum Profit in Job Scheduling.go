package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/maximum-profit-in-job-scheduling/
func init() {
	Solutions[1235] = func() {
		fmt.Println(`Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]`)
		fmt.Println(`Output:`, jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
		fmt.Println(`Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]`)
		fmt.Println(`Output:`, jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
	}
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	return 0
}

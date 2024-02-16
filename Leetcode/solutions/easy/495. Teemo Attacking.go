package easy

import "fmt"

// Reference: https://leetcode.com/problems/teemo-attacking/
func init() {
	Solutions[495] = func() {
		fmt.Println(`Input: timeSeries = [1,4], duration = 2`)
		fmt.Println(`Output:`, findPoisonedDuration([]int{1, 4}, 2))
		fmt.Println(`Input: timeSeries = [1,2], duration = 2`)
		fmt.Println(`Output:`, findPoisonedDuration([]int{1, 2}, 2))
	}
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	sum := 0
	for i := 1; i < n; i++ {
		sum += max(duration-(timeSeries[i]-timeSeries[i-1]), 0)
	}
	return n*duration - sum
}

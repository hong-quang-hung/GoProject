package hard

import "fmt"

// Reference: https://leetcode.com/problems/minimum-number-of-res-to-open-to-water-a-garden/
func init() {
	Solutions[1326] = func() {
		fmt.Println(`Input: n = 5, ranges = [3,4,1,1,0,0]`)
		fmt.Println(`Output:`, minTaps(5, []int{3, 4, 1, 1, 0, 0}))
		fmt.Println(`Input: n = 3, ranges = [0,0,0,0]`)
		fmt.Println(`Output:`, minTaps(3, []int{0, 0, 0, 0}))
	}
}

func minTaps(n int, ranges []int) int {
	maxReach := make([]int, n+1)
	for i := 0; i < len(ranges); i++ {
		start := max(0, i-ranges[i])
		end := min(n, i+ranges[i])
		maxReach[start] = max(maxReach[start], end)
	}

	res, currEnd, nextEnd := 0, 0, 0
	for i := 0; i <= n; i++ {
		if i > nextEnd {
			return -1
		}

		if i > currEnd {
			res++
			currEnd = nextEnd
		}
		nextEnd = max(nextEnd, maxReach[i])
	}
	return res
}

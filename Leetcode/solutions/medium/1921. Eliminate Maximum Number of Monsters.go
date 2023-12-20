package medium

import (
	"fmt"
	"math"
	"sort"
)

// Reference: https://leetcode.com/problems/eliminate-maximum-number-of-monsters/
func init() {
	Solutions[1921] = func() {
		fmt.Println(`Input: dist = [1,3,4], speed = [1,1,1]`)
		fmt.Println(`Output:`, eliminateMaximum([]int{1, 3, 4}, []int{1, 1, 1}))
		fmt.Println(`Input: dist = [1,1,2,3], speed = [1,1,1,1]`)
		fmt.Println(`Output:`, eliminateMaximum([]int{1, 1, 2, 3}, []int{1, 1, 1, 1}))
		fmt.Println(`Input: dist = [3,5,7,4,5], speed = [2,3,6,3,2]`)
		fmt.Println(`Output:`, eliminateMaximum([]int{3, 5, 7, 4, 5}, []int{2, 3, 6, 3, 2}))
	}
}

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	minutes := make([]int, n)
	for i := 0; i < n; i++ {
		minutes[i] = int(math.Ceil(float64(dist[i]) / float64(speed[i])))
	}

	sort.Ints(minutes)

	for i := 1; i < n; i++ {
		if minutes[i]-i <= 0 {
			return i
		}
	}
	return n
}

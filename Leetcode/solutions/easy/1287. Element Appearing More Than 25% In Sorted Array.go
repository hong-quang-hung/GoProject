package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/bus-routes/
func init() {
	Solutions[1287] = func() {
		fmt.Println(`Input: arr = [1,2,2,6,6,6,6,7,10]`)
		fmt.Println(`Output:`, findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
		fmt.Println(`Input: arr = [1,1]`)
		fmt.Println(`Output:`, findSpecialInteger([]int{1, 1}))
	}
}

func findSpecialInteger(arr []int) int {
	n := len(arr)
	target := n / 4
	nums := []int{arr[n/4], arr[n/2], arr[3*n/4]}
	for _, num := range nums {
		left := sort.Search(n, func(i int) bool { return arr[i] >= num })
		right := sort.Search(n, func(i int) bool { return arr[i] > num })
		if right-left > target {
			return num
		}
	}
	return -1
}

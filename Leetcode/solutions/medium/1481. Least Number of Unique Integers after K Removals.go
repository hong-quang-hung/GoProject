package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/
func init() {
	Solutions[1481] = func() {
		fmt.Println(`Input: arr = [5,5,4], k = 1`)
		fmt.Println(`Output:`, findLeastNumOfUniqueInts([]int{5, 5, 4}, 1))
		fmt.Println(`Input: arr = [4,3,1,1,3,3,2], k = 3`)
		fmt.Println(`Output:`, findLeastNumOfUniqueInts([]int{4, 3, 1, 1, 3, 3, 2}, 3))
	}
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := make(map[int]int)
	for _, a := range arr {
		m[a]++
	}

	n := len(m)
	nums := make([]int, n)
	i := 0
	for _, v := range m {
		nums[i] = v
		i++
	}

	sort.Ints(nums)

	i = 0
	for k > 0 {
		k -= nums[i]
		if k >= 0 {
			i++
		}
	}
	return n - i
}

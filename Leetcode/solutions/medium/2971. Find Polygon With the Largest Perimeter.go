package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/
func init() {
	Solutions[2971] = func() {
		fmt.Println(`Input: nums = [5,5,5]`)
		fmt.Println(`Output:`, largestPerimeter([]int{5, 5, 5}))
		fmt.Println(`Input: nums = [1,12,1,2,5,50,3]`)
		fmt.Println(`Output:`, largestPerimeter([]int{1, 12, 1, 2, 5, 50, 3}))
		fmt.Println(`Input: nums = [5,5,50]`)
		fmt.Println(`Output:`, largestPerimeter([]int{5, 5, 50}))
	}
}

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	res := -1
	prev := 0
	for _, num := range nums {
		if num < prev {
			res = num + prev
		}
		prev += num
	}
	return int64(res)
}

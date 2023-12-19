package medium

import (
	"fmt"
	"math"
	"sort"
)

// Reference: https://leetcode.com/problems/find-the-value-of-the-partition/
func init() {
	Solutions[2740] = func() {
		fmt.Println(`Input: nums = [1,3,2,4]`)
		fmt.Println(`Output:`, findValueOfPartition([]int{1, 3, 2, 4}))
		fmt.Println(`Input: nums = [100,1,10]`)
		fmt.Println(`Output:`, findValueOfPartition([]int{100, 1, 10}))
		fmt.Println(`Input: nums = [59,51,1,98,73]`)
		fmt.Println(`Output:`, findValueOfPartition([]int{59, 51, 1, 98, 73}))
	}
}

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	res := math.MaxInt
	for i := 1; i < len(nums); i++ {
		res = min(res, nums[i]-nums[i-1])
	}
	return res
}

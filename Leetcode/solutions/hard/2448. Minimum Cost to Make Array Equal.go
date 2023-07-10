package hard

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2448] = Leetcode_Min_Cost_III
}

// Reference: https://leetcode.com/problems/minimum-cost-to-make-array-equal/
func Leetcode_Min_Cost_III() {
	fmt.Println("Input: nums = [1,3,5,2], cost = [2,3,1,14]")
	fmt.Println("Output:", minCost_iii([]int{1, 3, 5, 2}, []int{2, 3, 1, 14}))
	fmt.Println("Input: nums = [2,2,2,2,2], cost = [4,2,8,1,3]")
	fmt.Println("Output:", minCost_iii([]int{2, 2, 2, 2, 2}, []int{4, 2, 8, 1, 3}))
}

func minCost_iii(nums []int, cost []int) int64 {
	n := len(nums)
	arr := make([][2]int, n)
	for i := 0; i < n; i++ {
		arr[i] = [2]int{nums[i], cost[i]}
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i][0] < arr[j][0] })

	prefixCost := make([]int64, n)
	prefixCost[0] = int64(arr[0][1])
	for i := 1; i < n; i++ {
		prefixCost[i] = prefixCost[i-1] + int64(arr[i][1])
	}

	res := int64(0)
	for i := 1; i < n; i++ {
		res += (int64(arr[i][0]) - int64(arr[0][0])) * int64(arr[i][1])
	}

	for i := 1; i < n; i++ {
		jump := int64(arr[i][0] - arr[i-1][0])
		res = min64(res, res-jump*prefixCost[n-1]+2*jump*int64(prefixCost[i-1]))
	}
	return res
}

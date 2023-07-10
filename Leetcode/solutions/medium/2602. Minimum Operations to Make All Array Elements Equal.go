package medium

import (
	"fmt"
	"math"
	"sort"
)

func init() {
	Solutions[2602] = Leetcode_Min_Operations_I
}

// Reference: https://leetcode.com/problems/minimum-operations-to-make-all-array-elements-equal/
func Leetcode_Min_Operations_I() {
	fmt.Println("Input: nums = [3,1,6,8], queries = [1,5]")
	fmt.Println("Output:", minOperations_i([]int{3, 1, 6, 8}, []int{1, 5}))
}

func minOperations_i(nums []int, queries []int) []int64 {
	var n int = len(nums)
	res := make([]int64, len(queries))
	sumIndex := make([]int64, n)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sumIndex[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		sumIndex[i] = sumIndex[i-1] + int64(nums[i])
	}

	for j, q := range queries {
		if q <= nums[0] || q > nums[n-1] {
			res[j] = int64(math.Abs(float64(int64(n)*int64(q) - int64(sumIndex[n-1]))))
		} else {
			index := findIndexMinOperations(nums, q)
			res[j] = int64(q)*int64(index) - 2*sumIndex[index-1] + sumIndex[n-1] - int64(q)*int64(n-index)
		}
	}

	return res
}

func findIndexMinOperations(nums []int, find int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < find {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

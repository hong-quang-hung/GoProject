package hard

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2659] = Leetcode_Count_Operations_To_Empty_Array
}

// Reference: https://leetcode.com/problems/make-array-empty/
func Leetcode_Count_Operations_To_Empty_Array() {
	fmt.Println("Input: nums = [3,4,-1]")
	fmt.Println("Output:", countOperationsToEmptyArray([]int{3, 4, -1}))
	fmt.Println("Input: nums = [1,2,4,3]")
	fmt.Println("Output:", countOperationsToEmptyArray([]int{1, 2, 4, 3}))
	fmt.Println("Input: nums = [1,2,3]")
	fmt.Println("Output:", countOperationsToEmptyArray([]int{1, 2, 3}))
}

func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)
	idx := make([][2]int, n)
	for i := 0; i < n; i++ {
		idx[i] = [2]int{nums[i], i}
	}

	sort.Slice(idx, func(i, j int) bool { return idx[i][0] < idx[j][0] })
	res := int64(n)
	for i := 1; i < n; i++ {
		if idx[i][1] < idx[i-1][1] {
			res += int64(n - i)
		}
	}
	return res
}

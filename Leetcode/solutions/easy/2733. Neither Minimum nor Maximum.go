package easy

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2733] = Leetcode_Find_Non_Min_Or_Max
}

// Reference: https://leetcode.com/problems/neither-minimum-nor-maximum/
func Leetcode_Find_Non_Min_Or_Max() {
	fmt.Println("Input: nums = [3,2,1,4]")
	fmt.Println("Output:", findNonMinOrMax([]int{3, 2, 1, 4}))
	fmt.Println("Input: nums = [1,2]")
	fmt.Println("Output:", findNonMinOrMax([]int{1, 2}))
	fmt.Println("Input: nums = [2,1,3]")
	fmt.Println("Output:", findNonMinOrMax([]int{2, 1, 3}))
}

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	sort.Ints(nums)
	return nums[1]
}

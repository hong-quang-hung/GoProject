package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/max-number-of-k-sum-pairs/
func Leetcode_Max_Operations() {
	fmt.Println("Input: nums = [1,2,3,4], k = 5")
	fmt.Println("Output:", maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println("Input: nums = [3,1,3,4,3], k = 6")
	fmt.Println("Output:", maxOperations([]int{3, 1, 3, 4, 3}, 6))
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	res := 0
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]+nums[j] == k {
			i++
			j--
			res++
		} else if nums[i]+nums[j] > k {
			j--
		} else {
			i++
		}
	}
	return res
}

package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
func init() {
	Solutions[1498] = func() {
		fmt.Println("Input: nums = [3,5,6,7], target = 9")
		fmt.Println("Output:", numSubseq([]int{3, 5, 6, 7}, 9))
		fmt.Println("Input: nums = [3,3,6,8], target = 10")
		fmt.Println("Output:", numSubseq([]int{3, 3, 6, 8}, 10))
		fmt.Println("Input: nums = [2,3,3,4,6,7], target = 12")
		fmt.Println("Output:", numSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
	}
}

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	res, n := 0, len(nums)
	power := make([]int, n)
	power[0] = 1
	for i := 1; i < n; i++ {
		power[i] = (power[i-1] * 2) % mod
	}

	for i := 0; i < n && nums[i]*2 <= target; i++ {
		mid := sort.SearchInts(nums, target-nums[i])
		if mid == n || nums[mid]+nums[i] > target {
			mid--
		}

		for mid+1 < n && nums[mid] == nums[mid+1] {
			mid++
		}

		res += power[mid-i]
		res %= mod
	}
	return res
}

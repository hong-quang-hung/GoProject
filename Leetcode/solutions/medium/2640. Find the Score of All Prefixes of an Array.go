package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/find-the-width-of-columns-of-a-grid/
func Leetcode_Find_Prefix_Score() {
	fmt.Println("Input: nums = [2,3,7,5,10]")
	fmt.Println("Output:", findPrefixScore(utils.S2SliceInt("[2,3,7,5,10]")))
	fmt.Println("Input: nums = [1,1,2,4,8,16]")
	fmt.Println("Output:", findPrefixScore(utils.S2SliceInt("[1,1,2,4,8,16]")))
}

func findPrefixScore(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)
	maxNums := nums[0]
	ans[0] = int64(nums[0]) + int64(maxNums)
	for i := 1; i < n; i++ {
		maxNums = max(maxNums, nums[i])
		ans[i] += ans[i-1] + int64(nums[i]) + int64(maxNums)
	}
	return ans
}

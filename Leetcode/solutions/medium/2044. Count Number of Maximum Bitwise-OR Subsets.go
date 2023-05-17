package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
func Leetcode_Count_Max_Or_Subsets() {
	fmt.Println("Input: nums = [3,1]")
	fmt.Println("Output:", countMaxOrSubsets([]int{3, 1}))
	fmt.Println("Input: nums = [2,2,2]")
	fmt.Println("Output:", countMaxOrSubsets([]int{2, 2, 2}))
	fmt.Println("Input: nums = [10,8,4]")
	fmt.Println("Output:", countMaxOrSubsets([]int{3, 2, 1, 5}))
}

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	pre := make([]int64, n+1)
	pre[0] = 0
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] | int64(nums[i])
	}

	fmt.Println(pre)
	return 0
}

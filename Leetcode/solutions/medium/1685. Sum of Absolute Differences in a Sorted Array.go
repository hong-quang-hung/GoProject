package medium

import "fmt"

// Reference: https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
func init() {
	Solutions[1685] = func() {
		fmt.Println("Input: nums = [2,3,5]")
		fmt.Println("Output:", getSumAbsoluteDifferences([]int{2, 3, 5}))
		fmt.Println("Input: nums = [1,4,6,8,10]")
		fmt.Println("Output:", getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10}))
	}
}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	prefix := make([]int, n)
	prefix[0] = 0
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + nums[i] - nums[i-1]
	}

	fmt.Println(prefix)

	for i := 0; i < n-1; i++ {

	}
	return res
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
func init() {
	Solutions[1685] = func() {
		fmt.Println(`Input: nums = [2,3,5]`)
		fmt.Println(`Output:`, getSumAbsoluteDifferences([]int{2, 3, 5}))
		fmt.Println(`Input: nums = [1,4,6,8,10]`)
		fmt.Println(`Output:`, getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10}))
	}
}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	res := make([]int, n)
	leftSum := 0
	for i := 0; i < n; i++ {
		rightSum := sum - leftSum - nums[i]
		leftCount, rightCount := i, n-1-i

		leftTotal := leftCount*nums[i] - leftSum
		rightTotal := rightSum - rightCount*nums[i]

		res[i] = leftTotal + rightTotal
		leftSum += nums[i]
	}
	return res
}

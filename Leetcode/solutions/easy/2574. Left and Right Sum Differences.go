package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/left-and-right-sum-differences/
func init() {
	Solutions[2574] = func() {
		fmt.Println("Input: nums = [10,4,8,3]")
		fmt.Println("Output:", leftRigthDifference(S2SliceInt("[10,4,8,3]")))
	}
}

func leftRigthDifference(nums []int) []int {
	leftSum, rightSum := 0, 0
	for _, num := range nums {
		rightSum += num
	}

	ans := make([]int, len(nums))
	for i, num := range nums {
		rightSum -= num
		ans[i] = int(math.Abs(float64(leftSum - rightSum)))
		leftSum += num
	}
	return ans
}

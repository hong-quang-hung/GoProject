package easy

import (
	"fmt"
	"math"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/left-and-right-sum-differences/
func Leetcode_Left_Rigth_Difference() {
	fmt.Println("Input: nums = [10,4,8,3]")
	fmt.Println("Output:", leftRigthDifference(utils.S2SliceInt("[10,4,8,3]")))
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

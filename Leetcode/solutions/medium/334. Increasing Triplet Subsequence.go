package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/increasing-triplet-subsequence/
func Leetcode_Increasing_Triplet() {
	fmt.Println("Input: nums = [1,2,3,4,5]")
	fmt.Println("Output:", increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println("Input: nums = [0,4,2,1,0,-1,-3]")
	fmt.Println("Output:", increasingTriplet([]int{0, 4, 2, 1, 0, -1, -3}))
	fmt.Println("Input: nums = [2,1,5,0,4,6]")
	fmt.Println("Output:", increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	fmt.Println("Input: nums = [20,100,10,12,5,13]")
	fmt.Println("Output:", increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

func increasingTriplet(nums []int) bool {
	val1, val2 := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num > val1 && num > val2 {
			return true
		}

		if num <= val1 {
			val1 = num
		} else if num <= val2 {
			val2 = num
		}
	}
	return false
}

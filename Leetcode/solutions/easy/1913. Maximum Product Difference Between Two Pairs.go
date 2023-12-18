package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/course-schedule-ii/
func init() {
	Solutions[1913] = func() {
		fmt.Println("Input: nums = [5,6,2,7,4]")
		fmt.Println("Output:", maxProductDifference([]int{5, 6, 2, 7, 4}))
		fmt.Println("Input: nums = [4,2,5,9,7,4,8]")
		fmt.Println("Output:", maxProductDifference([]int{4, 2, 5, 9, 7, 4, 8}))
	}
}

func maxProductDifference(nums []int) int {
	max1, max2 := 0, 0
	min1, min2 := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if max1 < num {
			max2, max1 = max1, num
		} else if max2 < num {
			max2 = num
		}
		if min1 > num {
			min2, min1 = min1, num
		} else if min2 > num {
			min2 = num
		}
	}
	return max1*max2 - min1*min2
}

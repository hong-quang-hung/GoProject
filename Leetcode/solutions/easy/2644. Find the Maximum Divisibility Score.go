package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/find-the-maximum-divisibility-score/
func init() {
	Solutions[2644] = func() {
		fmt.Println(`Input: nums = [4,7,9,3,9], divisors = [5,2,3]`)
		fmt.Println(`Output:`, maxDivScore([]int{4, 7, 9, 3, 9}, []int{5, 2, 3}))
		fmt.Println(`Input: nums = [20,14,21,10], divisors = [5,7,5]`)
		fmt.Println(`Output:`, maxDivScore([]int{20, 14, 21, 10}, []int{5, 7, 5}))
		fmt.Println(`Input: nums = [12], divisors = [10,16]`)
		fmt.Println(`Output:`, maxDivScore([]int{12}, []int{10, 16}))
	}
}

func maxDivScore(nums []int, divisors []int) int {
	res, maxCount := math.MaxInt, 0
	for i := 0; i < len(divisors); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j]%divisors[i] == 0 {
				count++
			}
		}

		if maxCount < count {
			res = divisors[i]
			maxCount = count
		} else if maxCount == count {
			if res > divisors[i] {
				res = divisors[i]
			}
		}
	}
	return res
}

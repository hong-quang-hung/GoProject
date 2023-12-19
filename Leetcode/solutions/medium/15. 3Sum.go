package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/3sum/
func init() {
	Solutions[15] = func() {
		fmt.Println(`Input: nums = [-1,0,1,2,-1,-4]`)
		fmt.Println(`Output:`, threeSum([]int{-1, 0, 1, 2, -1, -4}))
		fmt.Println(`Input: nums = [0,1,1]`)
		fmt.Println(`Output:`, threeSum([]int{0, 1, 1}))
	}
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}

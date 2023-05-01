package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/3sum-closest/
func Leetcode_Tree_Sum_Closest() {
	fmt.Println("Input: nums = [-1,2,1,-4], target = 1")
	fmt.Println("Output:", threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println("Input: nums = [0,0,0], target = 1")
	fmt.Println("Output:", threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println("Input: nums = [4,0,5,-5,3,3,0,-4,-5], target = -2")
	fmt.Println("Output:", threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
	fmt.Println("Input: nums = [1,1,1,0], target = 100")
	fmt.Println("Output:", threeSumClosest([]int{1, 1, 1, 0}, 100))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return res
			}
		}
	}
	return res
}

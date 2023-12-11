package medium

import "fmt"

// Reference: https://leetcode.com/problems/arithmetic-subarrays/
func init() {
	Solutions[1630] = func() {
		fmt.Println("Input: nums = [4,6,5,9,3,7], l = [0,0,2], r = [2,3,5]")
		fmt.Println("Output:", checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
		fmt.Println("Input: nums = [-12,-9,-3,-12,-6,15,20,-25,-20,-15,-10], l = [0,1,6,4,8,7], r = [4,4,9,7,9,10]")
		fmt.Println("Output:", checkArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
	}
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(l)
	res := make([]bool, len(l))
	for i := 0; i < n; i++ {
		res[i] = checkArithmeticSubarraysCheck(nums[l[i] : r[i]+1])
	}
	return res
}

func checkArithmeticSubarraysCheck(nums []int) bool {
	min := nums[0]
	max := nums[0]
	sum := 0
	var len = len(nums)
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	if max == min {
		return true
	}
	if (max-min)%(len-1) > 0 {
		return false
	}
	if (max+min)*len != 2*sum {
		return false
	}

	step := (max - min) / (len - 1)
	sum = 0
	for i, num := range nums {
		if ((num - min) % step) > 0 {
			return false
		}
		sum += i*i - (num-min)/step*(num-min)/step
	}
	return sum == 0
}

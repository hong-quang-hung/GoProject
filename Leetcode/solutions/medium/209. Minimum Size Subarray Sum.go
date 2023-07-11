package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-size-subarray-sum/
func init() {
	Solutions[209] = func() {
		fmt.Println("Input: target = 7, nums = [2,3,1,2,4,3]")
		fmt.Println("Output:", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
		fmt.Println("Input: target = 4, nums = [1,4,4]")
		fmt.Println("Output:", minSubArrayLen(4, []int{1, 4, 4}))
		fmt.Println("Input: target = 15, nums = [1,2,3,4,5]")
		fmt.Println("Output:", minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
	}
}

func minSubArrayLen(target int, nums []int) int {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	res, left, right := 100001, 0, 1
	for right < len(prefix) {
		if prefix[right]-prefix[left] < target {
			right++
		} else {
			if right-left < res {
				res = right - left
			}
			left++
		}
	}

	if res == 100001 {
		return 0
	}
	return res
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
func init() {
	Solutions[1658] = func() {
		fmt.Println("Input: nums = [1,1,4,2,3], x = 5")
		fmt.Println("Output:", minOperations([]int{1, 1, 4, 2, 3}, 5))
		fmt.Println("Input: nums = [5,6,7,8,9], x = 4")
		fmt.Println("Output:", minOperations([]int{5, 6, 7, 8, 9}, 4))
		fmt.Println("Input: nums = [3,2,20,1,1,3], x = 10")
		fmt.Println("Output:", minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
	}
}

func minOperations(nums []int, x int) int {
	n := len(nums)
	target := -x
	for _, num := range nums {
		target += num
	}

	if target == 0 {
		return n
	}

	res, sum, left := 0, 0, 0
	for right := 0; right < n; right++ {
		sum += nums[right]
		for left <= right && sum > target {
			sum -= nums[left]
			left++
		}

		if sum == target {
			res = max(res, right-left+1)
		}
	}

	if res != 0 {
		return n - res
	}
	return -1
}

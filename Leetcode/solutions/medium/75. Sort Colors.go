package medium

import "fmt"

// Reference: https://leetcode.com/problems/sort-colors/
func init() {
	Solutions[75] = func() {
		var nums []int
		fmt.Println("Input: nums = [2,0,2,1,1,0]")
		nums = []int{2, 0, 2, 1, 1, 0}
		sortColors(nums)
		fmt.Println("Output:", nums)
		fmt.Println("Input: nums = [2,0,1]")
		nums = []int{2, 0, 1}
		sortColors(nums)
		fmt.Println("Output:", nums)
	}
}

func sortColors(nums []int) {
	count := [3]int{}
	for _, num := range nums {
		count[num]++
	}

	idx := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < count[i]; j++ {
			nums[idx] = i
			idx++
		}
	}
}

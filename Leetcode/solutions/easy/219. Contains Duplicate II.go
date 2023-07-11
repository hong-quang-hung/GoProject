package easy

import "fmt"

// Reference: https://leetcode.com/problems/contains-duplicate-ii/
func init() {
	Solutions[219] = func() {
		fmt.Println("Input: nums = [1,2,3,1,2,3], k = 2")
		fmt.Println("Output:", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	distinct := make(map[int]int)
	for i, num := range nums {
		if d, c := distinct[num]; c && i-d <= k {
			return true
		} else {
			distinct[num] = i
		}
	}
	return false
}

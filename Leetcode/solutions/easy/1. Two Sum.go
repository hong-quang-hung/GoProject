package easy

import "fmt"

// Reference: https://leetcode.com/problems/two-sum/
func init() {
	Solutions[1] = func() {
		fmt.Println(`Input: nums = [2,7,11,15], target = 9`)
		fmt.Println(`Output:`, twoSum([]int{2, 7, 11, 15}, 9))
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if d, v := m[target-num]; v {
			return []int{d, i}
		}
		m[num] = i
	}
	return nil
}

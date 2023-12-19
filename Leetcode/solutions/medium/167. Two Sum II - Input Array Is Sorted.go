package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func init() {
	Solutions[167] = func() {
		fmt.Println(`Input: numbers = [2,7,11,15], target = 9`)
		fmt.Println(`Output:`, twoSumII([]int{2, 7, 11, 15}, 9))
		fmt.Println(`Input: numbers = [2,3,4], target = 6`)
		fmt.Println(`Output:`, twoSumII([]int{2, 3, 4}, 6))
		fmt.Println(`Input: numbers = [-1,0], target = -1`)
		fmt.Println(`Output:`, twoSumII([]int{-1, 0}, -1))
	}
}

func twoSumII(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		newTarget := target - numbers[i]
		j := sort.Search(n-i-1, func(x int) bool {
			return numbers[i+x+1] >= newTarget
		})

		nextIdx := i + j + 1
		if nextIdx < n && numbers[i+j+1] == newTarget {
			return []int{i + 1, nextIdx + 1}
		}
	}
	return []int{-1, -1}
}

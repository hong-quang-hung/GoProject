package hard

import "fmt"

// Reference: https://leetcode.com/problems/maximum-score-of-a-good-subarray/
func init() {
	Solutions[1793] = func() {
		fmt.Println("Input: nums = [1,4,3,7,4,5], k = 3")
		fmt.Println("Output:", maximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
		fmt.Println("Input: nums = [5,5,4,5,4,1,1,1], k = 0")
		fmt.Println("Output:", maximumScore([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0))
	}
}

func maximumScore(nums []int, k int) int {
	stack, res := [][2]int{}, 0
	for i, num := range nums {
		j := i
		for len(stack) > 0 && stack[len(stack)-1][1] >= num {
			idx, v := stack[len(stack)-1][0], stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			res = max(res, (i-idx+1)*v)
			j = idx
		}
		stack = append(stack, [2]int{j, num})
	}

	fmt.Println(stack)
	return res
}

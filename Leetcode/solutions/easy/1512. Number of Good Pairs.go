package easy

import "fmt"

// Reference: https://leetcode.com/problems/number-of-good-pairs/
func init() {
	Solutions[1512] = func() {
		fmt.Println(`Input: nums = [1,2,3,1,1,3]`)
		fmt.Println(`Output:`, numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
		fmt.Println(`Input: nums = [1,1,1,1]`)
		fmt.Println(`Output:`, numIdenticalPairs([]int{1, 1, 1, 1}))
	}
}

func numIdenticalPairs(nums []int) int {
	m := make([]int, 101)
	res := 0
	for _, num := range nums {
		res += m[num]
		m[num]++
	}
	return res
}

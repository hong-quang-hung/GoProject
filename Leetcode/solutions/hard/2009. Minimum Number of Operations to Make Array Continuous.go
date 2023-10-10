package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/
func init() {
	Solutions[2009] = func() {
		fmt.Println("Input: nums = [4,2,5,3]")
		fmt.Println("Output:", minOperations([]int{4, 2, 5, 3}))
		fmt.Println("Input: nums = [1,2,3,5,6]")
		fmt.Println("Output:", minOperations([]int{1, 2, 3, 5, 6}))
		fmt.Println("Input: nums = [41,33,29,33,35,26,47,24,18,28]")
		fmt.Println("Output:", minOperations([]int{41, 33, 29, 33, 35, 26, 47, 24, 18, 28}))
	}
}

func minOperations(nums []int) int {
	m := make(map[int]interface{})
	for _, num := range nums {
		m[num] = struct{}{}
	}

	distinct := []int{}
	for k := range m {
		distinct = append(distinct, k)
	}

	sort.Ints(distinct)

	n := len(nums)
	res := n
	for i := range distinct {
		find := distinct[i] + n - 1
		j := sort.Search(len(distinct), func(i int) bool { return distinct[i] > find })
		res = min(res, n-j+i)
	}
	return res
}

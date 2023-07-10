package hard

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2551] = Leetcode_Put_Marbles
}

// Reference: https://leetcode.com/problems/put-marbles-in-bags/
func Leetcode_Put_Marbles() {
	fmt.Println("Input: weights = [1,3,5,1], k = 2")
	fmt.Println("Output:", putMarbles([]int{1, 3, 5, 1}, 2))
	fmt.Println("Input: weights = [1, 3], k = 2")
	fmt.Println("Output:", putMarbles([]int{1, 3}, 2))
}

func putMarbles(weights []int, k int) int64 {
	lastIdx := len(weights) - 1
	for i := 0; i < lastIdx; i++ {
		weights[i] = weights[i] + weights[i+1]
	}

	weights = weights[:lastIdx]
	sort.Slice(weights, func(i, j int) bool {
		return weights[i] > weights[j]
	})

	res := 0
	for i := 0; i < k-1; i++ {
		res += weights[i] - weights[lastIdx-i-1]
	}
	return int64(res)
}

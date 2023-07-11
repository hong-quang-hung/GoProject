package easy

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
func init() {
	Solutions[1502] = func() {
		fmt.Println("Input: arr = [3,5,1]")
		fmt.Println("Output:", canMakeArithmeticProgression([]int{3, 5, 1}))
		fmt.Println("Input: arr = [1,2,4]")
		fmt.Println("Output:", canMakeArithmeticProgression([]int{1, 2, 4}))
	}
}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

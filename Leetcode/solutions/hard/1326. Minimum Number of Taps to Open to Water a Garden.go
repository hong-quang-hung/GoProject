package hard

import "fmt"

// Reference: https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/
func init() {
	Solutions[1326] = func() {
		fmt.Println("Input: n = 5, ranges = [3,4,1,1,0,0]")
		fmt.Println("Output:", minTaps(5, []int{3, 4, 1, 1, 0, 0}))
		fmt.Println("Input: n = 3, ranges = [0,0,0,0]")
		fmt.Println("Output:", minTaps(3, []int{0, 0, 0, 0}))
	}
}

func minTaps(n int, ranges []int) int {
	return 0
}

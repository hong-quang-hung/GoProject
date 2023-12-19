package medium

import "fmt"

// Reference: https://leetcode.com/problems/max-chunks-to-make-sorted/
func init() {
	Solutions[769] = func() {
		fmt.Println(`Input: arr = [4,3,2,1,0]`)
		fmt.Println(`Output:`, maxChunksToSorted([]int{4, 3, 2, 1, 0}))
		fmt.Println(`Input: arr = [1,0,2,3,4]`)
		fmt.Println(`Output:`, maxChunksToSorted([]int{1, 0, 2, 3, 4}))
	}
}

func maxChunksToSorted(arr []int) int {
	monotonic := []int{}
	for _, a := range arr {
		maxChunk := a
		for len(monotonic) > 0 && monotonic[len(monotonic)-1] > a {
			pop := monotonic[len(monotonic)-1]
			monotonic = monotonic[:len(monotonic)-1]
			maxChunk = max(maxChunk, pop)
		}
		monotonic = append(monotonic, maxChunk)
	}
	return len(monotonic)
}

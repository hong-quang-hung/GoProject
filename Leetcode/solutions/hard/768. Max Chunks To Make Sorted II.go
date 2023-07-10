package hard

import "fmt"

func init() {
	Solutions[768] = Leetcode_Max_Chunks_To_Sorted_II
}

// Reference: https://leetcode.com/problems/max-chunks-to-make-sorted-ii/
func Leetcode_Max_Chunks_To_Sorted_II() {
	fmt.Println("Input: arr = [4,2,2,1,1]")
	fmt.Println("Output:", maxChunksToSorted_ii([]int{4, 2, 2, 1, 1}))
	fmt.Println("Input: arr = [0,0,1,1,1]")
	fmt.Println("Output:", maxChunksToSorted_ii([]int{0, 0, 1, 1, 1}))
}

func maxChunksToSorted_ii(arr []int) int {
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

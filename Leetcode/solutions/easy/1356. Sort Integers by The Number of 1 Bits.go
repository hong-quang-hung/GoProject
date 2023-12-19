package easy

import (
	"fmt"
	"math/bits"
	"sort"
)

// Reference: https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
func init() {
	Solutions[1356] = func() {
		fmt.Println(`Input: arr = [0,1,2,3,4,5,6,7,8]`)
		fmt.Println(`Output:`, sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
		fmt.Println(`Input: arr = [1024,512,256,128,64,32,16,8,4,2,1]`)
		fmt.Println(`Output:`, sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	}
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		return bi < bj || bi == bj && arr[i] < arr[j]
	})
	return arr
}

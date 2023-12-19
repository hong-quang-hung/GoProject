package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-length-of-pair-chain/
func init() {
	Solutions[646] = func() {
		fmt.Println(`Input: pairs = [[1,2],[2,3],[3,4]]`)
		fmt.Println(`Output:`, findLongestChain(S2SoSliceInt(`[[1,2],[2,3],[3,4]]`)))
		fmt.Println(`Input: pairs = [[1,2],[7,8],[4,5]]`)
		fmt.Println(`Output:`, findLongestChain(S2SoSliceInt(`[[1,2],[7,8],[4,5]]`)))
	}
}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	res := 1
	right := pairs[0][1]
	for i := 0; i < len(pairs); i++ {
		if right < pairs[i][0] {
			res++
			right = pairs[i][1]
		}
	}
	return res
}

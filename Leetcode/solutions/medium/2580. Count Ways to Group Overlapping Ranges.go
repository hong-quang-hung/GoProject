package medium

import (
	"fmt"
	"sort"
)

func init() {
	Solutions[2580] = Leetcode_Count_Ways
}

// Reference: https://leetcode.com/problems/count-ways-to-group-overlapping-ranges/
func Leetcode_Count_Ways() {
	fmt.Println("Input: ranges = [[6,10],[5,15]]")
	fmt.Println("Output:", countWays(S2SoSliceInt("[[6,10],[5,15]]")))
}

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	overlapping, end := 1, ranges[0][1]
	for i := 1; i < len(ranges); i++ {
		if end < ranges[i][0] {
			overlapping++
		}
		end = max(end, ranges[i][1])
	}

	res := int64(1)
	for i := overlapping; i > 0; i-- {
		res <<= 1
		res %= mod
	}
	return int(res)
}

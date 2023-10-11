package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/number-of-flowers-in-full-bloom/
func init() {
	Solutions[2251] = func() {
		fmt.Println("Input: flowers = [[1,6],[3,7],[9,12],[4,13]], people = [2,3,7,11]")
		fmt.Println("Output:", fullBloomFlowers(S2SoSliceInt("[[1,6],[3,7],[9,12],[4,13]]"), []int{2, 3, 7, 11}))
		fmt.Println("Input: flowers = [[1,10],[3,3]], people = [3,3,2]")
		fmt.Println("Output:", fullBloomFlowers(S2SoSliceInt("[[1,10],[3,3]]"), []int{3, 3, 2}))
	}
}

func fullBloomFlowers(flowers [][]int, people []int) []int {
	sort.Slice(flowers, func(i, j int) bool {
		return flowers[i][0] < flowers[j][0]
	})

	fmt.Println(flowers)

	res := make([]int, len(people))
	for i, p := range people {
		j := sort.Search(len(flowers), func(i int) bool { return flowers[i][0] >= p })
		res[i] = j
		if j < len(flowers) && flowers[j][0] == p {
			res[i]++
		}
	}
	return res
}

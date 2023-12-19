package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/number-of-flowers-in-full-bloom/
func init() {
	Solutions[2251] = func() {
		fmt.Println(`Input: flowers = [[1,6],[3,7],[9,12],[4,13]], people = [2,3,7,11]`)
		fmt.Println(`Output:`, fullBloomFlowers(S2SoSliceInt(`[[1,6],[3,7],[9,12],[4,13]]`), []int{2, 3, 7, 11}))
		fmt.Println(`Input: flowers = [[1,10],[3,3]], people = [3,3,2]`)
		fmt.Println(`Output:`, fullBloomFlowers(S2SoSliceInt(`[[1,10],[3,3]]`), []int{3, 3, 2}))
	}
}

func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	starts, ends := make([]int, n), make([]int, n)
	for i, flower := range flowers {
		starts[i], ends[i] = flower[0], flower[1]+1
	}

	sort.Ints(starts)
	sort.Ints(ends)

	res := make([]int, 0)
	for _, person := range people {
		i := sort.Search(n, func(i int) bool { return starts[i] > person })
		j := sort.Search(n, func(i int) bool { return ends[i] > person })
		res = append(res, i-j)
	}
	return res
}

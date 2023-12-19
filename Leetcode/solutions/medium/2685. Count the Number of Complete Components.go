package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-the-number-of-complete-components/
func init() {
	Solutions[2685] = func() {
		fmt.Println(`Input: n = 6, edges = [[0,1],[0,2],[1,2],[3,4]]`)
		fmt.Println(`Output:`, countCompleteComponents(6, S2SoSliceInt(`[[0,1],[0,2],[1,2],[3,4]]`)))
		fmt.Println(`Input: n = 6, edges = [[0,1],[0,2],[1,2],[3,4],[3,5]]`)
		fmt.Println(`Output:`, countCompleteComponents(6, S2SoSliceInt(`[[0,1],[0,2],[1,2],[3,4],[3,5]]`)))
	}
}

func countCompleteComponents(n int, edges [][]int) int {
	union, counter := NewUnionFind(n), make([]int, n)
	for _, edge := range edges {
		union.UnionSet(edge[0], edge[1])
		counter[edge[0]]++
		counter[edge[1]]++
	}

	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[union.Find(i)] = true
	}

	for i := 0; i < n; i++ {
		find := union.Find(i)
		if m[find] && union.Count[find] != counter[i]+1 {
			delete(m, find)
		}
	}
	return len(m)
}

package medium

import (
	"fmt"
	"sort"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/smallest-string-with-swaps/
func Leetcode_Smallest_String_With_Swaps() {
	fmt.Println("Input: s = 'dcab', pairs = [[0,3],[1,2]]")
	fmt.Println("Output:", smallestStringWithSwaps("dcab", utils.S2SoSliceInt("[[0,3],[1,2]]")))
	fmt.Println("Input: s = 'dcab', pairs = [[0,3],[1,2],[0,2]]")
	fmt.Println("Output:", smallestStringWithSwaps("dcab", utils.S2SoSliceInt("[[0,3],[1,2],[0,2]]")))
	fmt.Println("Input: s = 'cba', pairs = [[0,1],[1,2]]")
	fmt.Println("Output:", smallestStringWithSwaps("cba", utils.S2SoSliceInt("[[0,1],[1,2]]")))
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	union := types.NewUnionFind(n)
	for _, pair := range pairs {
		union.UnionSet(pair[0], pair[1])
	}

	for i := 0; i < n; i++ {
		union.Parent[i] = union.Find(union.Parent[i])
	}

	m := make(map[int][]byte)
	for i := 0; i < n; i++ {
		list := m[union.Parent[i]]
		if list == nil {
			list = []byte{}
		}
		list = append(list, s[i])
		m[union.Parent[i]] = list
	}

	for k := range m {
		sort.Slice(m[k], func(i, j int) bool { return m[k][i] < m[k][j] })
	}

	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = m[union.Parent[i]][0]
		m[union.Parent[i]] = m[union.Parent[i]][1:]
	}
	return string(res)
}

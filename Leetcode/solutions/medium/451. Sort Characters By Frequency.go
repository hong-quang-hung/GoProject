package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/sort-characters-by-frequency/
func init() {
	Solutions[451] = func() {
		fmt.Println(`Input: s = "tree"`)
		fmt.Println(`Output:`, frequencySort("tree"))
		fmt.Println(`Input: s = "cccaaa"`)
		fmt.Println(`Output:`, frequencySort("cccaaa"))
		fmt.Println(`Input: s = "Aabb"`)
		fmt.Println(`Output:`, frequencySort("Aabb"))
	}
}

func frequencySort(s string) string {
	m := make([][2]int, 77)
	for _, ch := range s {
		key := int(ch - '.')
		m[key][0]++
		m[key][1] = key
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i][0] > m[j][0]
	})

	res := []rune{}
	for i := 0; i < 77; i++ {
		if m[i][0] == 0 {
			break
		}

		r := rune(m[i][1] + '.')
		for j := 0; j < m[i][0]; j++ {
			res = append(res, r)
		}
	}
	return string(res)
}

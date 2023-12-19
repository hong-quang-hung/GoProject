package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/binary-trees-with-factors/
func init() {
	Solutions[823] = func() {
		fmt.Println(`Input: arr = [2,4]`)
		fmt.Println(`Output:`, numFactoredBinaryTrees([]int{2, 4}))
		fmt.Println(`Input: arr = [2,4,5,10]`)
		fmt.Println(`Output:`, numFactoredBinaryTrees([]int{2, 4, 5, 10}))
		fmt.Println(`Input: arr = [2,4,8]`)
		fmt.Println(`Output:`, numFactoredBinaryTrees([]int{2, 4, 8}))
	}
}

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)

	m := make(map[int]int64)
	for _, num := range arr {
		m[num] = 1
	}

	for _, n1 := range arr {
		for _, n2 := range arr {
			if int64(n2*n2) > int64(n1) {
				break
			}

			if n1%n2 == 0 {
				if v, ok := m[n1/n2]; ok {
					temp := (m[n2] * v) % mod
					if n1/n2 == n2 {
						m[n1] = (m[n1] + temp) % mod
					} else {
						m[n1] = (m[n1] + temp*2) % mod
					}
				}
			}
		}
	}

	res := 0
	for _, v := range m {
		res = (res + int(v)) % mod
	}
	return res
}

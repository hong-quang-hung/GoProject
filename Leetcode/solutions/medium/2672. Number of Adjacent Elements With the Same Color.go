package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-adjacent-elements-with-the-same-color/
func init() {
	Solutions[2672] = func() {
		fmt.Println("Input: n = 4, queries = [[0,2],[1,2],[3,1],[1,1],[2,1]]")
		fmt.Println("Output:", colorTheArray(4, S2SoSliceInt("[[0,2],[1,2],[3,1],[1,1],[2,1]]")))
		fmt.Println("n = 1, queries = [[0,100000]]")
		fmt.Println("Output:", colorTheArray(1, S2SoSliceInt("[[0,100000]]")))
	}
}

func colorTheArray(n int, queries [][]int) []int {
	res := make([]int, len(queries))
	m := make([]int, n)
	count := 0
	for i, querie := range queries {
		old := m[querie[0]]
		new := querie[1]
		if old != 0 {
			if querie[0]-1 >= 0 && m[querie[0]-1] == old {
				count--
			}

			if querie[0]+1 < n && m[querie[0]+1] == old {
				count--
			}
		}

		m[querie[0]] = new
		if new != 0 {
			if querie[0]-1 >= 0 && m[querie[0]-1] == new {
				count++
			}

			if querie[0]+1 < n && m[querie[0]+1] == new {
				count++
			}
		}

		res[i] = count
	}
	return res
}

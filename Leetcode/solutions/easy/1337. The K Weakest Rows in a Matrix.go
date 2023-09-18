package easy

import "fmt"

// Reference: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
func init() {
	Solutions[1337] = func() {
		fmt.Println("Input: mat = [[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]], k = 3")
		fmt.Println("Output:", kWeakestRows(S2SoSliceInt("[[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]]"), 3))
	}
}

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	buckets := make([][]int, n+1)
	for row := 0; row < m; row++ {
		i, j := 0, n
		for i < j {
			h := i + (j-i)/2
			if mat[row][h] > 0 {
				i = h + 1
			} else {
				j = h
			}
		}
		buckets[i] = append(buckets[i], row)
	}

	res := make([]int, 0, k)
	for _, bucket := range buckets {
		for i := 0; i < len(bucket) && k > 0; i++ {
			res = append(res, bucket[i])
			k--
		}
	}
	return res
}

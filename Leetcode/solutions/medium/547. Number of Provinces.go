package medium

import "fmt"

// Reference: https://leetcode.com/problems/number-of-provinces/
func init() {
	Solutions[547] = func() {
		fmt.Println("Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]")
		fmt.Println("Output:", findCircleNum(S2SoSliceInt("[[1,1,0],[1,1,0],[0,0,1]]")))
		fmt.Println("Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]")
		fmt.Println("Output:", findCircleNum(S2SoSliceInt("[[1,0,0],[0,1,0],[0,0,1]]")))
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	dsu := NewUnionFind(n)
	res := n
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 && dsu.Find(i) != dsu.Find(j) {
				res--
				dsu.UnionSet(i, j)
			}
		}
	}
	return res
}

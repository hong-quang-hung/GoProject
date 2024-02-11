package hard

import "fmt"

// Reference: https://leetcode.com/problems/cherry-pickup-ii/
func init() {
	Solutions[1463] = func() {
		fmt.Println(`Input: grid = [[3,1,1],[2,5,1],[1,5,5],[2,1,1]]`)
		fmt.Println(`Output:`, cherryPickup(S2SoSliceInt("[[3,1,1],[2,5,1],[1,5,5],[2,1,1]]")))
		fmt.Println(`Input: grid = [[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]]`)
		fmt.Println(`Output:`, cherryPickup(S2SoSliceInt("[[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]]")))
	}
}

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for z := 0; z < n; z++ {
				dp[i][j][z] = -1
			}
		}
	}

	var f func(i, j, k int) int
	f = func(i, j, k int) int {
		if i < 0 || i >= j || j == n || k == m {
			return 0
		}

		if dp[k][i][j] != -1 {
			return dp[k][i][j]
		}

		cherry := 0
		currCherry := grid[k][i] + grid[k][j]
		for dir1 := -1; dir1 <= 1; dir1++ {
			for dir2 := -1; dir2 <= 1; dir2++ {
				cherry = max(cherry, currCherry+f(i+dir1, j+dir2, k+1))
			}
		}

		dp[k][i][j] = cherry
		return cherry
	}
	return f(0, n-1, 0)
}

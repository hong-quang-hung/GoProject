package medium

import "fmt"

// Reference: https://leetcode.com/problems/spiral-matrix-ii/
func init() {
	Solutions[59] = func() {
		fmt.Println(`Input: n = 3`)
		fmt.Println(`Output:`, generateMatrix(3))
		fmt.Println(`Input: n = 1`)
		fmt.Println(`Output:`, generateMatrix(1))
	}
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	left, right, top, bottom, count := 0, n-1, 0, n-1, 1
	for count <= n*n {
		for i := left; i < right+1; i++ {
			res[top][i] = count
			count++
		}

		top++
		for i := top; i < bottom+1; i++ {
			res[i][right] = count
			count++
		}

		right--
		if top <= bottom {
			for i := right; i > left-1; i-- {
				res[bottom][i] = count
				count++
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i > top-1; i-- {
				res[i][left] = count
				count++
			}
			left++
		}
	}
	return res
}

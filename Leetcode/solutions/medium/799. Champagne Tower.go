package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-string-chain/
func init() {
	Solutions[799] = func() {
		fmt.Println("Input: poured = 1, query_row = 1, query_glass = 1")
		fmt.Println("Output:", champagneTower(1, 1, 1))
		fmt.Println("Input: poured = 2, query_row = 1, query_glass = 1")
		fmt.Println("Output:", champagneTower(2, 1, 1))
		fmt.Println("Input: poured = 100000009, query_row = 33, query_glass = 17")
		fmt.Println("Output:", champagneTower(100000009, 33, 17))
	}
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	arr := [102][102]float64{}
	arr[0][0] = float64(poured)
	for r := 0; r <= query_row; r++ {
		for c := 0; c <= r; c++ {
			q := (arr[r][c] - 1.0) / 2.0
			if q > 0 {
				arr[r+1][c] += q
				arr[r+1][c+1] += q
			}
		}
	}

	if arr[query_row][query_glass] < 1 {
		return arr[query_row][query_glass]
	}
	return 1
}

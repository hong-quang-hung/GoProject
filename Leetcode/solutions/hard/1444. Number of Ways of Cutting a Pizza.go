package hard

import "fmt"

func init() {
	Solutions[1444] = Leetcode_Ways
}

// Reference: https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/
func Leetcode_Ways() {
	fmt.Println("Input: pizza = ['A..','AAA,'...'], k = 3")
	fmt.Println("Output:", ways([]string{"A..", "AAA", "..."}, 3))
	fmt.Println("Input: pizza = ['A..','AA.,'...'], k = 3")
	fmt.Println("Output:", ways([]string{"A..", "AA.", "..."}, 3))
}

func ways(pizza []string, k int) int {
	rows, cols := len(pizza), len(pizza[0])
	apples := make([][]int, rows+1)
	f := make([][]int, rows)

	for row := 0; row < rows; row++ {
		apples[row] = make([]int, cols+1)
		f[row] = make([]int, cols)
	}
	apples[rows] = make([]int, cols+1)

	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if pizza[row][col] == 'A' {
				apples[row][col] = 1
			}
			apples[row][col] += apples[row+1][col] + apples[row][col+1] - apples[row+1][col+1]
			if apples[row][col] > 0 {
				f[row][col] = 1
			}
		}
	}

	const mod = int(1e9 + 7)
	for remain := 1; remain < k; remain++ {
		g := make([][]int, rows)
		for row := 0; row < rows; row++ {
			g[row] = make([]int, cols)
		}

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				for next_row := row + 1; next_row < rows; next_row++ {
					if apples[row][col]-apples[next_row][col] > 0 {
						g[row][col] += f[next_row][col]
						g[row][col] %= mod
					}
				}

				for next_col := col + 1; next_col < cols; next_col++ {
					if apples[row][col]-apples[row][next_col] > 0 {
						g[row][col] += f[row][next_col]
						g[row][col] %= mod
					}
				}
			}
		}
		copy(f, g)
	}
	return f[0][0]
}

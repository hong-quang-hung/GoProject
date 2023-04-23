package easy

import "fmt"

// Reference: https://leetcode.com/problems/pascals-triangle/
func Leetcode_Generate() {
	fmt.Println("Input: numRows = 5")
	fmt.Println("Output:", generate(5))
	fmt.Println("Input: numRows = 1")
	fmt.Println("Output:", generate(1))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if i == 0 || j == 0 || j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j] + res[i-1][j-1]
			}
		}
	}
	return res
}

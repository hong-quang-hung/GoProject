package easy

import (
	"fmt"
	"strconv"
)

func init() {
	Solutions[2639] = Leetcode_Max_Value_Of_Coins
}

// Reference: https://leetcode.com/problems/find-the-score-of-all-prefixes-of-an-array/
func Leetcode_Max_Value_Of_Coins() {
	fmt.Println("Input: grid = [[1],[22],[333]]")
	fmt.Println("Output:", findColumnWidth(S2SoSliceInt("[[1],[22],[333]]")))
	fmt.Println("Input: grid = [[-15,1,3],[15,7,12],[5,6,-2]]")
	fmt.Println("Output:", findColumnWidth(S2SoSliceInt("[[-15,1,3],[15,7,12],[5,6,-2]]")))
}

func findColumnWidth(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	res := make([]int, m)
	for i := 0; i < m; i++ {
		maxLenght := 0
		for j := 0; j < n; j++ {
			maxLenght = max(maxLenght, len(strconv.Itoa(grid[j][i])))
		}
		res[i] = maxLenght
	}
	return res
}

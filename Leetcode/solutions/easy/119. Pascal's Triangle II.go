package easy

import "fmt"

func init() {
	Solutions[119] = Leetcode_Get_Row
}

// Reference: https://leetcode.com/problems/pascals-triangle-ii/
func Leetcode_Get_Row() {
	fmt.Println("Input: numRows = 3")
	fmt.Println("Output:", getRow(3))
	fmt.Println("Input: numRows = 1")
	fmt.Println("Output:", getRow(1))
}

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		old := make([]int, rowIndex+1)
		copy(old, res)
		for j := 0; j <= i; j++ {
			if i == 0 || j == 0 || j == i {
				res[j] = 1
			} else {
				res[j] = old[j] + old[j-1]
			}
		}
	}
	return res
}

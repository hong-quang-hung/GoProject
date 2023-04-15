package hard

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/
func Leetcode_Max_Value_Of_Coins() {
	fmt.Println("Input: piles = [[1,100,3],[7,8,9]], k = 2")
	fmt.Println("Output:", maxValueOfCoins(utils.S2SoSliceInt("[[1,100,3],[7,8,9]]"), 2))
	fmt.Println("Input: piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7")
	fmt.Println("Output:", maxValueOfCoins(utils.S2SoSliceInt("[[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]]"), 7))
}

func maxValueOfCoins(piles [][]int, k int) int {
	return k
}

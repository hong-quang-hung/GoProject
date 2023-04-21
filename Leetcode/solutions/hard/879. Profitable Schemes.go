package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/profitable-schemes/
func Leetcode_Profitable_Schemes() {
	fmt.Println("Input: n = 5, minProfit = 3, group = [2,2], profit = [2,3]")
	fmt.Println("Output:", profitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
	fmt.Println("Input: n = 10, minProfit = 5, group = [2,3,5], profit = [6,7,8]")
	fmt.Println("Output:", profitableSchemes(5, 3, []int{2, 3, 5}, []int{6, 7, 8}))
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	return 0
}

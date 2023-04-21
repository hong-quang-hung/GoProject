package hard

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/minimize-the-total-price-of-the-trips/
func Leetcode_Minimum_Total_Price() {
	fmt.Println("Input: n = 4, edges = [[0,1],[1,2],[1,3]], price = [2,2,10,6], trips = [[0,3],[2,1],[2,3]]")
	fmt.Println("Output:", minimumTotalPrice(4, utils.S2SoSliceInt("[[0,1],[1,2],[1,3]]"), []int{2, 2, 10, 6}, utils.S2SoSliceInt("[[0,3],[2,1],[2,3]]")))
	fmt.Println("Input: n = 2, edges = [[0,1]], price = [2,2], trips = [[0,0]]")
	fmt.Println("Output:", minimumTotalPrice(2, utils.S2SoSliceInt("[[0,1]]"), []int{2, 2}, utils.S2SoSliceInt("[[0,0]]")))
}

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	return 0
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/cheapest-flights-within-k-stops/
func init() {
	Solutions[787] = func() {
		fmt.Println(`Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1`)
		fmt.Println(`Output:`, findCheapestPrice(4, S2SoSliceInt("[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]]"), 0, 3, 1))
		fmt.Println(`Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1`)
		fmt.Println(`Output:`, findCheapestPrice(3, S2SoSliceInt("[[0,1,100],[1,2,100],[0,2,500]]"), 0, 2, 1))
		fmt.Println(`Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0`)
		fmt.Println(`Output:`, findCheapestPrice(3, S2SoSliceInt("[[0,1,100],[1,2,100],[0,2,500]]"), 0, 2, 0))
	}
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	return 0
}

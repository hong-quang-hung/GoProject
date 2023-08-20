package hard

import "fmt"

// Reference: https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/
func init() {
	Solutions[1203] = func() {
		fmt.Println("Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]")
		fmt.Println("Output:", sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, S2SoSliceInt("[[],[6],[5],[6],[3,6],[],[],[]]")))
		fmt.Println("Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]")
		fmt.Println("Output:", sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, S2SoSliceInt("[[],[6],[5],[6],[3],[],[4],[]]")))
	}
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	return nil
}

package medium

import "leetcode.com/Leetcode/utils"

func min(a, b int) int {
	return utils.MinInt(a, b)
}

func max(a, b int) int {
	return utils.MaxInt(a, b)
}

func max64(a, b int64) int64 {
	return utils.MaxInt64(a, b)
}

func abs(x int) int {
	return utils.AbsInt(x)
}

func manhattanDistance(p1, p2 []int) int {
	return utils.ManhattanDistanceInt(p1, p2)
}

func insert(arr []int, index int, value int) []int {
	return utils.InsertInt(arr, index, value)
}

func gcd(a, b int) int {
	return utils.GcdInt(a, b)
}

func gcd64(a, b int64) int64 {
	return utils.GcdInt64(a, b)
}

func compareSliceInt(a, b []int) bool {
	return utils.CompareSliceInt(a, b)
}

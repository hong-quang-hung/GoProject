package medium

import "leetcode.com/Leetcode/utils"

func min(a, b int) int {
	return utils.MinInt(a, b)
}

func max(a, b int) int {
	return utils.MaxInt(a, b)
}

func manhattanDistance(p1, p2 []int) int {
	return utils.ManhattanDistanceInt(p1, p2)
}

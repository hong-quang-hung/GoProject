package solutions

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ManhattanDistance(p1, p2 []int) int {
	return int(math.Abs(float64(p1[0]-p2[0]))) + int(math.Abs(float64(p1[1]-p2[1])))
}

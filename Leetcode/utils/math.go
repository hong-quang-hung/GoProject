package utils

import "math"

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ManhattanDistanceInt(p1, p2 []int) int {
	return int(math.Abs(float64(p1[0]-p2[0]))) + int(math.Abs(float64(p1[1]-p2[1])))
}

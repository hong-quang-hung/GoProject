package utils

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

func MaxInt64(a, b int64) int64 {
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

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func AbsInt(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func ManhattanDistanceInt(p1, p2 []int) int {
	return AbsInt(p1[0]-p2[0]) + AbsInt(p1[0]-p2[0])
}

func GcdInt(a, b int) int {
	if a == 0 {
		return b
	}
	return GcdInt(b%a, a)
}

func GcdInt64(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return GcdInt64(b%a, a)
}

// Compare two slice which same length
// Constraints: len(a) == len(b)
func CompareSliceInt(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

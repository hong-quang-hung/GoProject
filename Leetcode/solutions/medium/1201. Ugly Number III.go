package medium

import "fmt"

// Reference: https://leetcode.com/problems/ugly-number-iii/
func init() {
	Solutions[1201] = func() {
		fmt.Println("Input: n = 3, a = 2, b = 3, c = 5")
		fmt.Println("Output:", nthUglyNumber_iii(3, 2, 3, 5))
		fmt.Println("Input: n = 4, a = 2, b = 3, c = 4")
		fmt.Println("Output:", nthUglyNumber_iii(4, 2, 3, 4))
		fmt.Println("Input: n = 5, a = 2, b = 11, c = 13")
		fmt.Println("Output:", nthUglyNumber_iii(5, 2, 11, 13))
	}
}

func nthUglyNumber_iii(n int, a int, b int, c int) int {
	low, hight := int64(1), int64(2e9)
	ab, ac, bc := lcm(int64(a), int64(b)), lcm(int64(a), int64(c)), lcm(int64(b), int64(c))
	abc := lcm(ab, int64(c))
	for low < hight {
		mid := low + (hight-low)/2
		count := mid/int64(a) + mid/int64(b) + mid/int64(c) - mid/ab - mid/ac - mid/bc + mid/abc
		if count >= int64(n) {
			hight = mid
		} else {
			low = mid + 1
		}
	}
	return int(low)
}

func lcm(a, b int64) int64 {
	return a * b / gcd64(a, b)
}

package hard

import "fmt"

// Reference: https://leetcode.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/
func init() {
	Solutions[1420] = func() {
		fmt.Println("Input: n = 2, m = 3, k = 1")
		fmt.Println("Output:", numOfArrays(2, 3, 1))
		fmt.Println("Input: n = 5, m = 2, k = 3")
		fmt.Println("Output:", numOfArrays(5, 2, 3))
		fmt.Println("Input: n = 9, m = 1, k = 1")
		fmt.Println("Output:", numOfArrays(9, 1, 1))
	}
}

func numOfArrays(n int, m int, k int) int {
	dp := make(map[string]int)
	var f func(id, length, lar int) int
	f = func(id, length, lar int) int {
		if id == n {
			if length == k {
				return 1
			} else {
				return 0
			}
		}

		key := fmt.Sprintf("%d-%d-%d", id, length, lar)
		if v, ok := dp[key]; ok {
			return v
		}

		res := 0
		for i := 1; i <= m; i++ {
			if i > lar {
				res += f(id+1, length+1, i)
			} else {
				res += f(id+1, length, lar)
			}
			res = res % mod
		}

		dp[key] = res
		return res
	}
	return f(0, 0, 0)
}

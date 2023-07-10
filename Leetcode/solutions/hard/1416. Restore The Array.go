package hard

import (
	"fmt"
	"strconv"
)

func init() {
	Solutions[1416] = Leetcode_Number_Of_Arrays
}

// Reference: https://leetcode.com/problems/restore-the-array/
func Leetcode_Number_Of_Arrays() {
	fmt.Println("Input: s = '1000', k = 10000")
	fmt.Println("Output:", numberOfArrays("1000", 10000))
	fmt.Println("Input: s = '1000', k = 10")
	fmt.Println("Output:", numberOfArrays("1000", 10))
	fmt.Println("Input: s = '1317', k = 2000")
	fmt.Println("Output:", numberOfArrays("1317", 2000))
}

func numberOfArrays(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)
	return numberOfArraysDFS(s, k, dp, 0, n)
}

func numberOfArraysDFS(s string, k int, dp []int, start int, n int) int {
	if dp[start] != 0 {
		return dp[start]
	}

	if start == n {
		return 1
	}

	if s[start] == '0' {
		return 0
	}

	count := 0
	for i := start; i < n; i++ {
		curNumber, _ := strconv.ParseInt(s[start:i+1], 10, 64)
		if curNumber > int64(k) {
			break
		}

		count = (count + numberOfArraysDFS(s, k, dp, i+1, n)) % mod
	}

	dp[start] = count
	return count
}

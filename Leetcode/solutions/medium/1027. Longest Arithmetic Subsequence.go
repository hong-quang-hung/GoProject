package medium

import "fmt"

func init() {
	Solutions[1027] = Leetcode_Longest_Arith_Seq_Length
}

// Reference: https://leetcode.com/problems/longest-arithmetic-subsequence/
func Leetcode_Longest_Arith_Seq_Length() {
	fmt.Println("Input: nums = [3,6,9,12]")
	fmt.Println("Output:", longestArithSeqLength([]int{3, 6, 9, 12}))
	fmt.Println("Input: nums = [9,4,7,2,10]")
	fmt.Println("Output:", longestArithSeqLength([]int{9, 4, 7, 2, 10}))
	fmt.Println("Input: nums = [20,1,15,3,10,5,8]")
	fmt.Println("Output:", longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8}))
}

func longestArithSeqLength(nums []int) int {
	n := len(nums)
	dp := make([][1001]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 1001; j++ {
			dp[i][j] = 1
		}
	}

	res := 2
	dp[1][nums[1]-nums[0]+500] = 2
	for i := 2; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			diff := nums[i] - nums[j] + 500
			dp[i][diff] = max(dp[i][diff], dp[j][diff]+1)
			res = max(res, dp[i][diff])
		}
	}
	return res
}

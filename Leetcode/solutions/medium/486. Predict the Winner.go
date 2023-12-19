package medium

import "fmt"

// Reference: https://leetcode.com/problems/predict-the-winner/
func init() {
	Solutions[486] = func() {
		fmt.Println(`Input: nums = [1,5,2]`)
		fmt.Println(`Output:`, predictTheWinner([]int{1, 5, 2}))
		fmt.Println(`Input: nums = [1,5,233,7]`)
		fmt.Println(`Output:`, predictTheWinner([]int{1, 5, 233, 7}))
	}
}

func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
	}

	first, second := predictTheWinnerSolved(nums, dp, 0, 0, true)
	return first >= second
}

func predictTheWinnerSolved(nums []int, dp [][][]int, left int, right int, first bool) (int, int) {
	if len(nums) == left+right+1 {
		if first {
			return nums[left], 0
		} else {
			return 0, nums[left]
		}
	}

	if dp[left][right] != nil {
		return dp[left][right][0], dp[left][right][1]
	}

	first1, second1 := predictTheWinnerSolved(nums, dp, left+1, right, !first)
	first2, second2 := predictTheWinnerSolved(nums, dp, left, right+1, !first)

	if first {
		first1 = first1 + nums[left]
		first2 = first2 + nums[len(nums)-1-right]

		if first1 > first2 || (first1 == first2 && second1 <= second2) {
			dp[left][right] = []int{first1, second1}
		} else {
			dp[left][right] = []int{first2, second2}
		}
	} else {
		second1 = second1 + nums[left]
		second2 = second2 + nums[len(nums)-1-right]

		if second1 > second2 || (second1 == second2 && first1 < first2) {
			dp[left][right] = []int{first1, second1}
		} else {
			dp[left][right] = []int{first2, second2}
		}
	}
	return dp[left][right][0], dp[left][right][1]
}

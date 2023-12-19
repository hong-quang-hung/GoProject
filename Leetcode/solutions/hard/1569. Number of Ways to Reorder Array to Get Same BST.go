package hard

import "fmt"

// Reference: https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/
func init() {
	Solutions[1569] = func() {
		fmt.Println(`Input: nums = [2,1,3]`)
		fmt.Println(`Output:`, numOfWays([]int{2, 1, 3}))
		fmt.Println(`Input: nums = [3,4,5,1,2]`)
		fmt.Println(`Output:`, numOfWays([]int{3, 4, 5, 1, 2}))
		fmt.Println(`Input: nums = [1, 2, 3]`)
		fmt.Println(`Output:`, numOfWays([]int{1, 2, 3}))
	}
}

func numOfWays(nums []int) int {
	table := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		table[i] = make([]int, i+1)
		table[i][0] = 1
		table[i][i] = 1
		for j := 1; j < i; j++ {
			table[i][j] = (table[i-1][j] + table[i-1][j-1]) % mod
		}
	}
	return numOfWaysSolved(nums, table) - 1
}

func numOfWaysSolved(nums []int, table [][]int) int {
	n := len(nums)
	if n <= 2 {
		return 1
	}

	root := nums[0]
	left, right := []int{}, []int{}
	for i := 1; i < n; i++ {
		if nums[i] < root {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	lc := numOfWaysSolved(left, table) % mod
	rc := numOfWaysSolved(right, table) % mod
	return (((lc * rc) % mod) * table[n-1][len(left)]) % mod
}

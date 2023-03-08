package solutions

import (
	"math"

	U "leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/as-far-from-land-as-possible/
func maxDistance(grid [][]int) int {
	var max int = -1

	n, max := len(grid)-1, -1
	queue := [][2]int{}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		land := queue[0]
		queue = queue[1:]
		point := grid[land[0]][land[1]]
		col, row := land[0], land[1]

		if col > 0 && grid[col-1][row] == 0 {
			grid[col-1][row] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col - 1, row})
		}

		if col < n && grid[col+1][row] == 0 {
			grid[col+1][row] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col + 1, row})
		}

		if row > 0 && grid[col][row-1] == 0 {
			grid[col][row-1] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col, row - 1})
		}

		if row < n && grid[col][row+1] == 0 {
			grid[col][row+1] = point + 1
			if max < point {
				max = point
			}
			queue = append(queue, [2]int{col, row + 1})
		}
	}
	return max
}

// Reference: https://leetcode.com/problems/sort-an-array/
func sortArray(nums []int) []int {
	U.CountingSort(nums)
	return nums
}

// Given a string s, find the length of the longest substring without repeating characters.
// Reference: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	chars := [128]int{}
	left, right, res := 0, 0, 0
	for right < len(s) {
		char := s[right]
		index := chars[char] - 1
		if (index != 0 || char == s[0]) && index >= left && index < right {
			left = index + 1
		}
		if res < (right - left + 1) {
			res = right - left + 1
		}
		chars[char] = right + 1
		right++
	}
	return res
}

// Reference: https://leetcode.com/problems/count-total-number-of-colored-cells/
func coloredCells(n int) int64 {
	return int64(2*n*n - 2*n + 1)
}

// Reference: https://leetcode.com/problems/koko-eating-bananas/
func minEatingSpeed(piles []int, h int) int {
	var left int = 1
	var right int = 0
	for _, p := range piles {
		if right < p {
			right = p
		}
	}
	var res int = right
	for left <= right {
		var hours int64 = 0
		var mid = left + (right-left)/2
		for _, p := range piles {
			hours += int64(math.Ceil(float64(p) / float64(mid)))
		}
		if hours <= int64(h) {
			if res > mid {
				res = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

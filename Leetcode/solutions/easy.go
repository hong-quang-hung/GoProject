package solutions

import (
	T "leetcode.com/Leetcode/types"
)

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// Reference: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if d, v := m[target-num]; v {
			return []int{d, i}
		}
		m[num] = i
	}
	return nil
}

// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
// Reference: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *T.TreeNode) int {
	if root == nil {
		return 0
	}
	var l int = maxDepth(root.Left)
	var r int = maxDepth(root.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

// Given an integer x, return true if x is a palindrome, and false otherwise.
// Reference: https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	r := 0
	t := x
	for x > 0 {
		r = (r * 10) + (x % 10)
		x = x / 10

	}
	return t == r
}

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
func maxProfit(prices []int) int {
	maxPrices, minPrices := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrices > prices[i] {
			minPrices = prices[i]
		} else if maxPrices < prices[i]-minPrices {
			maxPrices = prices[i] - minPrices
		}
	}
	return maxPrices
}

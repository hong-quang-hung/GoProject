package solutions

import T "leetcode.com/Leetcode/types"

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

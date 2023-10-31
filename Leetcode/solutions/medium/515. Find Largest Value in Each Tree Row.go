package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/find-largest-value-in-each-tree-row/
func init() {
	Solutions[515] = func() {
		fmt.Println("Input: root = [1,3,2,5,3,null,9]")
		fmt.Println("Output:", largestValues(S2TreeNode("[1,3,2,5,3,null,9]")))
		fmt.Println("Input: root = [1,2,3]")
		fmt.Println("Output:", largestValues(S2TreeNode("[1,2,3]")))
	}
}

func largestValues(root *TreeNode) []int {
	q := []*TreeNode{}
	if root != nil {
		q = append(q, root)
	}

	res := []int{}
	for len(q) > 0 {
		next := []*TreeNode{}
		maxVal := math.MinInt
		for _, node := range q {
			maxVal = max(maxVal, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, maxVal)
		q = next
	}
	return res
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/
func init() {
	Solutions[2385] = func() {
		fmt.Println(`Input: root = [1,5,3,null,4,10,6,9,2], start = 3`)
		fmt.Println(`Output:`, amountOfTime(S2TreeNode("[1,5,3,null,4,10,6,9,2]"), 3))
		fmt.Println(`Input: root = [1], start = 1`)
		fmt.Println(`Output:`, amountOfTime(S2TreeNode("[1]"), 1))
	}
}

func amountOfTime(root *TreeNode, start int) int {
	res := 0
	var f func(node *TreeNode, val int) (int, bool)
	f = func(node *TreeNode, val int) (int, bool) {
		if node == nil {
			return 0, false
		}

		left, hasLeft := f(node.Left, val)
		right, hasRight := f(node.Right, val)

		if node.Val == start {
			res = max(res, max(left, right))
			return 1, true
		}

		if hasLeft {
			res = max(res, right+left)
			return left + 1, true

		} else if hasRight {
			res = max(res, right+left)
			return right + 1, true

		} else {
			return max(left, right) + 1, false
		}
	}

	f(root, start)
	return res
}

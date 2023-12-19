package medium

import "fmt"

// Reference: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func init() {
	Solutions[230] = func() {
		fmt.Println(`Input: root = [3,1,4,null,2], k = 1`)
		fmt.Println(`Output:`, kthSmallest(S2TreeNode(`[3,1,4,null,2]`), 1))
		fmt.Println(`Input: root = [5,3,6,2,4,null,null,1], k = 3`)
		fmt.Println(`Output:`, kthSmallest(S2TreeNode(`[5,3,6,2,4,null,null,1]`), 3))
		fmt.Println(`Input: root = [1,null,2], k = 2`)
		fmt.Println(`Output:`, kthSmallest(S2TreeNode(`[1,null,2]`), 2))
	}
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	return kthSmallestDFS(root, k, &count)
}

func kthSmallestDFS(root *TreeNode, k int, count *int) int {
	if root.Left == nil && root.Right == nil {
		(*count)++
		if k == *count {
			return root.Val
		}
		return -1
	}

	res := -1
	if root.Left != nil {
		res = kthSmallestDFS(root.Left, k, count)
	}

	(*count)++
	if k == *count {
		return root.Val
	}

	if res == -1 && root.Right != nil {
		res = kthSmallestDFS(root.Right, k, count)
	}
	return res
}

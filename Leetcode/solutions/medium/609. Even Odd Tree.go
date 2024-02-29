package medium

import "fmt"

// Reference: https://leetcode.com/problems/even-odd-tree/
func init() {
	Solutions[609] = func() {
		fmt.Println(`Input: root = [1,10,4,3,null,7,9,12,8,6,null,null,2]`)
		fmt.Println(`Output:`, isEvenOddTree(S2TreeNode("[1,10,4,3,null,7,9,12,8,6,null,null,2]")))
		fmt.Println(`Input: root = [5,4,2,3,3,7]`)
		fmt.Println(`Output:`, isEvenOddTree(S2TreeNode("[5,4,2,3,3,7]")))
		fmt.Println(`Input: root = [5,9,1,3,5,7]`)
		fmt.Println(`Output:`, isEvenOddTree(S2TreeNode("[5,9,1,3,5,7]")))
	}
}

func isEvenOddTree(root *TreeNode) bool {
	odd := true
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := []*TreeNode{}
		var prev int
		if odd {
			prev = 0
		} else {
			prev = 1000001
		}

		for _, node := range queue {
			if (node.Val%2 == 1) != odd {
				return false
			}

			if node.Val == prev || ((node.Val > prev) != odd) {
				return false
			}

			prev = node.Val
			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}

		odd = !odd
		queue = temp
		temp = nil
	}
	return true
}

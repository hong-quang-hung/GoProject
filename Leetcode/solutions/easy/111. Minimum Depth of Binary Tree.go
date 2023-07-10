package easy

import "fmt"

// Reference: https://leetcode.com/problems/minimum-depth-of-binary-tree/
func Leetcode_Min_Depth() {
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output:", minDepth(S2TreeNode("[3,9,20,null,null,15,7]")))
	fmt.Println("Input: root = [2,null,3,null,4,null,5,null,6]")
	fmt.Println("Output:", minDepth(S2TreeNode("[2,null,3,null,4,null,5,null,6]")))
	fmt.Println("Input: root = [1,2,3,4,5]")
	fmt.Println("Output:", minDepth(S2TreeNode("[1,2,3,4,5]")))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		temp := []*TreeNode{}
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				return res
			}

			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}

		res++
		q = temp
	}
	return res
}

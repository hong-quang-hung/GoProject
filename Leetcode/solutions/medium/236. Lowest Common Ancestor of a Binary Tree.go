package medium

import "fmt"

// Reference: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
func Leetcode_Lowest_Common_Ancestor() {
	fmt.Println("Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1")
	fmt.Println("Output:", lowestCommonAncestor(S2TreeNode("[3,5,1,6,2,0,8,null,null,7,4]"), nil, nil))
	fmt.Println("Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4")
	fmt.Println("Output:", lowestCommonAncestor(S2TreeNode("[3,5,1,6,2,0,8,null,null,7,4]"), nil, nil))
	fmt.Println("Input: root = [1,2], p = 1, q = 2")
	fmt.Println("Output:", lowestCommonAncestor(S2TreeNode("[1,2]"), nil, nil))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return nil
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/house-robber-iii/
func init() {
	Solutions[337] = func() {
		fmt.Println("Input: root = [3,2,3,null,3,null,1]")
		fmt.Println("Output:", rob_iii(S2TreeNode("[3,2,3,null,3,null,1]")))
		fmt.Println("Input: root = [3,4,5,1,3,null,1]")
		fmt.Println("Output:", rob_iii(S2TreeNode("[3,4,5,1,3,null,1]")))
	}
}

func rob_iii(root *TreeNode) int {
	first, second := rob_iii_solved(root)
	return max(first, second)
}

func rob_iii_solved(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	f1, f2 := rob_iii_solved(root.Left)
	s1, s2 := rob_iii_solved(root.Right)
	return root.Val + f2 + s2, max(s1, s2) + max(f1, f2)
}

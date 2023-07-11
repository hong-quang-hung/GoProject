package medium

import "fmt"

// Reference: https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
func init() {
	Solutions[863] = func() {
		root := S2TreeNode("[3,5,1,6,2,0,8,null,null,7,4]")
		target := root.Left
		fmt.Println("Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2")
		fmt.Println("Output:", distanceK(root, target, 2))

		root = S2TreeNode("[1]")
		target = root
		fmt.Println("Input: root = [1], target = 1, k = 3")
		fmt.Println("Output:", distanceK(root, target, 3))
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	return nil
}

package easy

import "fmt"

// Reference: https://leetcode.com/problems/average-of-levels-in-binary-tree/
func init() {
	Solutions[637] = func() {
		fmt.Println("Input: root = [3,9,20,null,null,15,7]")
		fmt.Println("Output:", averageOfLevels(S2TreeNode("[3,9,20,null,null,15,7]")))
		fmt.Println("Input: root = [3,9,20,15,7]")
		fmt.Println("Output:", averageOfLevels(S2TreeNode("[3,9,20,15,7]")))
	}
}

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := []*TreeNode{}
		sum := 0
		for _, node := range queue {
			sum += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		res = append(res, float64(sum)/float64(len(queue)))
		queue = next
	}
	return res
}

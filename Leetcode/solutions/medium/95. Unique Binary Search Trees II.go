package medium

import "fmt"

// Reference: https://leetcode.com/problems/unique-binary-search-trees-ii/
func init() {
	Solutions[95] = func() {
		fmt.Println("Input: n = 3")
		fmt.Println("Output:", SAny(generateTrees(3)))
		fmt.Println("Input: n = 1")
		fmt.Println("Output:", SAny(generateTrees(1)))
	}
}

func generateTrees(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = []*TreeNode{}
	}

	dp[0] = append(dp[0], nil)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			z := i - 1
			for _, left := range dp[j-1] {
				for _, right := range dp[z] {
					dp[i] = append(dp[i], &TreeNode{Val: j, Left: left, Right: generateTreesClone(right, j)})
				}
			}
		}
	}
	return dp[n]
}

func generateTreesClone(node *TreeNode, offset int) *TreeNode {
	if node == nil {
		return nil
	}

	clonedNode := &TreeNode{Val: node.Val + offset}
	clonedNode.Left = generateTreesClone(node.Left, offset)
	clonedNode.Right = generateTreesClone(node.Right, offset)
	return clonedNode
}

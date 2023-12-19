package medium

import "fmt"

// Reference: https://leetcode.com/problems/validate-binary-tree-nodes/
func init() {
	Solutions[1361] = func() {
		fmt.Println(`Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]`)
		fmt.Println(`Output:`, validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))
		fmt.Println(`Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]`)
		fmt.Println(`Output:`, validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}))
		fmt.Println(`Input: n = 2, leftChild = [1,0], rightChild = [-1,-1]`)
		fmt.Println(`Output:`, validateBinaryTreeNodes(2, []int{1, 0}, []int{-1, -1}))
	}
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	dsu := NewDSU(n)
	for i := 0; i < n; i++ {
		children := []int{leftChild[i], rightChild[i]}
		for _, child := range children {
			if child == -1 {
				continue
			}

			if child != dsu.Find(child) || !dsu.Join(i, child) {
				return false
			}
		}
	}
	return dsu.Components == 1
}

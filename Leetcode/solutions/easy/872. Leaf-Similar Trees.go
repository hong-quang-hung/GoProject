package easy

import (
	"fmt"
	"reflect"
)

func init() {
	Solutions[872] = Leetcode_Leaf_Similar
}

// Reference: https://leetcode.com/problems/leaf-similar-trees/
func Leetcode_Leaf_Similar() {
	fmt.Println("Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]")
	fmt.Println("Output:", leafSimilar(S2TreeNode("[3,5,1,6,2,9,8,null,null,7,4]"), S2TreeNode("[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]")))
	fmt.Println("Input: root1 = [1,2,3], root2 = [1,3,2]")
	fmt.Println("Output:", leafSimilar(S2TreeNode("[1,2,3]"), S2TreeNode("[1,3,2]")))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1, leaf2 := []int{}, []int{}
	leafSimilarDFS(root1, &leaf1)
	leafSimilarDFS(root2, &leaf2)
	return reflect.DeepEqual(leaf1, leaf2)
}

func leafSimilarDFS(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		(*arr) = append((*arr), root.Val)
	}

	leafSimilarDFS(root.Left, arr)
	leafSimilarDFS(root.Right, arr)
}

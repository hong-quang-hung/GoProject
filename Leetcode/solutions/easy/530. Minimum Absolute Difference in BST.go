package easy

import (
	"fmt"
	"math"
)

func init() {
	Solutions[530] = Leetcode_Get_Minimum_Difference
}

// Reference: https://leetcode.com/problems/minimum-absolute-difference-in-bst/
func Leetcode_Get_Minimum_Difference() {
	fmt.Println("Input: root = [4,2,6,1,3]")
	fmt.Println("Output:", getMinimumDifference(S2TreeNode("[4,2,6,1,3]")))
	fmt.Println("Input: root = [1,0,48,null,null,12,49]")
	fmt.Println("Output:", getMinimumDifference(S2TreeNode("[1,0,48,null,null,12,49]")))
}

func getMinimumDifference(root *TreeNode) int {
	res, preValue := math.MaxInt, -1
	getMinimumDifferenceDFS(root, &preValue, &res)
	return res
}

func getMinimumDifferenceDFS(root *TreeNode, preValue *int, minDiff *int) {
	if root == nil {
		return
	}

	getMinimumDifferenceDFS(root.Left, preValue, minDiff)

	if *preValue != -1 && root.Val-*preValue < *minDiff {
		*minDiff = root.Val - *preValue
	}

	*preValue = root.Val
	getMinimumDifferenceDFS(root.Right, preValue, minDiff)
}

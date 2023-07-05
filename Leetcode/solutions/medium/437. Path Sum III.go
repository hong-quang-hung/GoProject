package medium

import "fmt"

// Reference: https://leetcode.com/problems/path-sum-iii/
func Leetcode_Path_Sum_III() {
	fmt.Println("Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8")
	fmt.Println("Output:", pathSum(S2TreeNode("[10,5,-3,3,2,null,11,3,-2,null,1]"), 8))
	fmt.Println("Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22")
	fmt.Println("Output:", pathSum(S2TreeNode("[5,4,8,11,null,13,4,7,2,null,null,5,1]"), 22))
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	pathSumCount(root, targetSum, &res)
	return res
}

func pathSumCount(root *TreeNode, targetSum int, res *int) int {
	if root == nil {
		return 0
	}
	return 0
}

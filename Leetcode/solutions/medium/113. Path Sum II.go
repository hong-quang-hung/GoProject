package medium

import "fmt"

// Reference: https://leetcode.com/problems/path-sum-ii/
func init() {
	Solutions[113] = func() {
		fmt.Println(`Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22`)
		fmt.Println(`Output:`, pathSum(S2TreeNode(`[5,4,8,11,null,13,4,7,2,null,null,5,1]`), 22))
		fmt.Println(`Input: root = [1,2,3], targetSum = 5`)
		fmt.Println(`Output:`, pathSum(S2TreeNode(`[1,2,3]`), 5))
		fmt.Println(`Input: root = [-6,null,-3,-6,0,-6,-5,4,null,null,null,1,7], targetSum = -21`)
		fmt.Println(`Output:`, pathSum(S2TreeNode(`[-6,null,-3,-6,0,-6,-5,4,null,null,null,1,7]`), -21))
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	pathSumSolved(root, targetSum, []int{}, &res)
	return res
}

func pathSumSolved(root *TreeNode, targetSum int, curr []int, res *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			temp := make([]int, len(curr))
			copy(temp, curr)
			(*res) = append((*res), append(temp, root.Val))
		}
		return
	}

	curr = append(curr, root.Val)
	targetSum -= root.Val
	pathSumSolved(root.Left, targetSum, curr, res)
	pathSumSolved(root.Right, targetSum, curr, res)
}

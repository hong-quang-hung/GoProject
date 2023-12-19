package medium

import "fmt"

// Reference: https://leetcode.com/problems/path-sum-iii/
func init() {
	Solutions[437] = func() {
		fmt.Println(`Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8`)
		fmt.Println(`Output:`, pathSum_iii(S2TreeNode(`[10,5,-3,3,2,null,11,3,-2,null,1]`), 8))
		fmt.Println(`Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22`)
		fmt.Println(`Output:`, pathSum_iii(S2TreeNode(`[5,4,8,11,null,13,4,7,2,null,null,5,1]`), 22))
	}
}

func pathSum_iii(root *TreeNode, targetSum int) int {
	m := map[int]int{}
	m[0] = 1
	pathSumCount(root, targetSum, 0, m)
	return pathSumCount(root, targetSum, 0, m)
}

func pathSumCount(root *TreeNode, targetSum int, curr int, m map[int]int) int {
	if root == nil {
		return 0
	}

	curr += root.Val
	res := m[curr-targetSum]

	m[curr]++
	res += pathSumCount(root.Left, targetSum, curr, m)
	res += pathSumCount(root.Right, targetSum, curr, m)
	m[curr]--
	return res
}

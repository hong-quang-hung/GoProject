package easy

import "fmt"

// Reference: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
func init() {
	Solutions[108] = func() {
		fmt.Println(`Input: nums = [-10,-3,0,5,9]`)
		fmt.Println(`Output:`, STreeNode(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
		fmt.Println(`Input: nums = [1,3]`)
		fmt.Println(`Output:`, STreeNode(sortedArrayToBST([]int{1, 3})))
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	mid := (n - 1) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

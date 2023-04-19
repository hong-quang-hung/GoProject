package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
func Leetcode_Sorted_List_To_BST() {
	fmt.Println("Input: head = [-10,-3,0,5,9]")
	fmt.Print("Output:", utils.STreeNode(sortedListToBST(types.NewListNode(-10, -3, 0, 5, 9))))
	fmt.Println("Input: head = []")
	fmt.Print("Output:", utils.STreeNode(sortedListToBST(types.NewListNode())))
}

func sortedListToBST(head *types.ListNode) *types.TreeNode {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return createSortedListToBSTNode(values, 0, len(values)-1)
}

func createSortedListToBSTNode(values []int, start, end int) *types.TreeNode {
	if start < 0 || end >= len(values) || start > end {
		return nil
	}

	mid := (start + end) / 2
	return &types.TreeNode{Val: values[mid], Left: createSortedListToBSTNode(values, start, mid-1), Right: createSortedListToBSTNode(values, mid+1, end)}
}

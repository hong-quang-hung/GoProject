package medium

import "fmt"

// Reference: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
func init() {
	Solutions[109] = func() {
		fmt.Println("Input: head = [-10,-3,0,5,9]")
		fmt.Print("Output:", STreeNode(sortedListToBST(S2ListNode("[-10,-3,0,5,9]"))))
		fmt.Println("Input: head = []")
		fmt.Print("Output:", STreeNode(sortedListToBST(S2ListNode("[]"))))
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return createSortedListToBSTNode(values, 0, len(values)-1)
}

func createSortedListToBSTNode(values []int, start, end int) *TreeNode {
	if start < 0 || end >= len(values) || start > end {
		return nil
	}

	mid := (start + end) / 2
	return &TreeNode{Val: values[mid], Left: createSortedListToBSTNode(values, start, mid-1), Right: createSortedListToBSTNode(values, mid+1, end)}
}

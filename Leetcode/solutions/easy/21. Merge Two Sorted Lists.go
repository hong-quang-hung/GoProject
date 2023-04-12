package easy

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/merge-two-sorted-lists/
func Leetcode_Merge_Two_Lists() {
	fmt.Println("Input: list1 = [1,2,4], list2 = [1,3,4]")
	fmt.Println("Output:", mergeTwoLists(types.NewListNode(1, 2, 4), types.NewListNode(1, 3, 4)))
}

func mergeTwoLists(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {
	head := new(types.ListNode)
	newList := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newList.Next = &types.ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			newList.Next = &types.ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		newList = newList.Next
	}

	for list1 != nil {
		newList.Next = &types.ListNode{Val: list1.Val}
		list1 = list1.Next
		newList = newList.Next
	}

	for list2 != nil {
		newList.Next = &types.ListNode{Val: list2.Val}
		list2 = list2.Next
		newList = newList.Next
	}
	return head.Next
}

package easy

import "fmt"

// Reference: https://leetcode.com/problems/merge-two-sorted-lists/
func init() {
	Solutions[21] = func() {
		fmt.Println("Input: list1 = [1,2,4], list2 = [1,3,4]")
		fmt.Println("Output:", mergeTwoLists(S2ListNode("1,2,4"), S2ListNode("1,3,4")))
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	newList := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newList.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			newList.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		newList = newList.Next
	}

	for list1 != nil {
		newList.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		newList = newList.Next
	}

	for list2 != nil {
		newList.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		newList = newList.Next
	}
	return head.Next
}

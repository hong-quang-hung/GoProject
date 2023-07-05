package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/reverse-linked-list/
func Leetcode_Reverse_List() {
	fmt.Println("Input: head = [1,2,3,4,5]")
	fmt.Println("Output:", SListNode(reverseList(S2ListNode("[1,2,3,4,5]"))))
	fmt.Println("Input: head = [1,2]")
	fmt.Println("Output:", SListNode(reverseList(S2ListNode("[2,1]"))))
	fmt.Println("Input: head = []")
	fmt.Println("Output:", SListNode(reverseList(S2ListNode("[]"))))
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}

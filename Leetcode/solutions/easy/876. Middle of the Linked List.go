package easy

import "fmt"

// Reference: https://leetcode.com/problems/middle-of-the-linked-list/
func init() {
	Solutions[876] = func() {
		fmt.Println(`Input: head = [1,2,3,4,5]`)
		fmt.Println(`Output:`, SListNode(middleNode(S2ListNode("[1,2,3,4,5]"))))
		fmt.Println(`Input: head = [1,2,3,4,5,6]`)
		fmt.Println(`Output:`, SListNode(middleNode(S2ListNode("[1,2,3,4,5,6]"))))
	}
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
func Leetcode_Delete_Middle() {
	fmt.Println("Input: head = [1,3,4,7,1,2,6]")
	fmt.Println("Output:", deleteMiddle(S2ListNode("[1,3,4,7,1,2,6]")))
	fmt.Println("Input: head = [1,2,3,4]")
	fmt.Println("Output:", deleteMiddle(S2ListNode("[1,2,3,4]")))
	fmt.Println("Input: head = [2,1]")
	fmt.Println("Output:", deleteMiddle(S2ListNode("[2,1]")))
	fmt.Println("Input: head = [1]")
	fmt.Println("Output:", deleteMiddle(S2ListNode("[1]")))
}

func deleteMiddle(head *ListNode) *ListNode {
	fast, slow, mid := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid = slow
		slow = slow.Next
	}

	if mid != slow {
		return nil
	}

	mid.Next = slow.Next
	return head
}

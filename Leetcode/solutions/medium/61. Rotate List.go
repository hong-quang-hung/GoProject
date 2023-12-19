package medium

import "fmt"

// Reference: https://leetcode.com/problems/rotate-list/
func init() {
	Solutions[61] = func() {
		fmt.Println(`Input: head = [1,2,3,4,5], k = 2`)
		fmt.Println(`Output:`, SListNode(rotateRight(S2ListNode(`[1,2,3,4,5]`), 2)))
		fmt.Println(`Input: head = [0,1,2], k = 4`)
		fmt.Println(`Output:`, SListNode(rotateRight(S2ListNode(`[0,1,2]`), 4)))
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	n := 1
	temp := head
	for temp.Next != nil {
		n++
		temp = temp.Next
	}

	temp.Next = head

	k = n - k%n
	for k > 0 {
		temp = temp.Next
		k--
	}

	head = temp.Next
	temp.Next = nil
	return head
}

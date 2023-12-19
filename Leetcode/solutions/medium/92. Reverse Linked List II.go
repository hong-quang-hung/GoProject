package medium

import "fmt"

// Reference: https://leetcode.com/problems/reverse-linked-list-ii/
func init() {
	Solutions[92] = func() {
		fmt.Println(`Input: head = [1,2,3,4,5], left = 2, right = 4`)
		fmt.Println(`Output:`, SListNode(reverseBetween(S2ListNode(`[1,2,3,4,5]`), 2, 4)))
		fmt.Println(`Input: head = [5], left = 1, right = 1`)
		fmt.Println(`Output:`, SListNode(reverseBetween(S2ListNode(`[	5]`), 1, 1)))
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var left_end *ListNode
	curr := head
	for i := 0; i < left-1; i++ {
		left_end = curr
		curr = curr.Next
	}

	middle_end := curr
	var prev *ListNode
	for i := left; i < right+1; i++ {
		curr_next := curr.Next
		curr.Next = prev
		prev = curr
		curr = curr_next
	}

	middle_start := prev
	right_start := curr

	middle_end.Next = right_start
	if left_end != nil {
		left_end.Next = middle_start
		return head
	}
	return middle_start
}

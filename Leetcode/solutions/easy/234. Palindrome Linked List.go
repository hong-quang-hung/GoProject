package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/palindrome-linked-list/
func init() {
	Solutions[234] = func() {
		fmt.Println("Input: head = [1,2,2,1]")
		fmt.Println("Output:", isPalindromeII(S2ListNode("[1,2,2,1]")))
		fmt.Println("Input: head = [1,2]")
		fmt.Println("Output:", isPalindromeII(S2ListNode("[1,2]")))
	}
}

func isPalindromeII(head *ListNode) bool {
	var reverse, slow, fast *ListNode
	reverse, slow, fast = nil, head, head
	for fast != nil && fast.Next != nil {
		next := slow
		slow = slow.Next
		fast = fast.Next.Next
		next.Next = reverse
		reverse = next
	}

	if fast != nil {
		slow = slow.Next
	}

	for reverse != nil && slow != nil {
		if reverse.Val != slow.Val {
			return false
		}

		reverse = reverse.Next
		slow = slow.Next
	}
	return true
}

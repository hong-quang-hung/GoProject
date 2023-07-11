package medium

import "fmt"

// Reference: https://leetcode.com/problems/swap-nodes-in-pairs/
func init() {
	Solutions[24] = func() {
		fmt.Println("Input: head = [1,2,3,4]")
		fmt.Println("Output:", swapPairs(S2ListNode("[1,2,3,4,5]")))
		fmt.Println("Input: head = [1]")
		fmt.Println("Output:", swapPairs(S2ListNode("[1]")))
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head
	}
	return head
}

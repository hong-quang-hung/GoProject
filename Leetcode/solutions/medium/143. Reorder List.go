package medium

import "fmt"

// Reference: https://leetcode.com/problems/reorder-list/
func init() {
	Solutions[143] = func() {
		var head *ListNode

		fmt.Println(`Input: head = [1,2,3,4]`)
		head = S2ListNode("[1,2,3,4]")
		reorderList(head)
		fmt.Println(`Output:`, SListNode(head))

		fmt.Println(`Input: head = [1,2,3,4,5]`)
		head = S2ListNode("[1,2,3,4,5]")
		reorderList(head)
		fmt.Println(`Output:`, SListNode(head))
	}
}

func reorderList(head *ListNode) {
	fast, slow, mid := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		mid = slow
	}

	curr := mid.Next
	mid.Next = nil

	var reverse *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = reverse
		reverse = curr
		curr = next
	}

	dump := head
	for reverse != nil {
		nextDump := dump.Next
		nextReverse := reverse.Next
		dump.Next = reverse
		reverse.Next = nextDump
		dump = nextDump
		reverse = nextReverse
	}
}

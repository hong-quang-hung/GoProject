package medium

import "fmt"

// Reference: https://leetcode.com/problems/partition-list/
func init() {
	Solutions[86] = func() {
		fmt.Println("Input: head = [1,4,3,2,5,2], x = 3")
		fmt.Println("Output:", SListNode(partition(S2ListNode("[1,4,3,2,5,2]"), 3)))
		fmt.Println("Input: head = [2,1], x = 2")
		fmt.Println("Output:", SListNode(partition(S2ListNode("[2,1]"), 2)))
	}
}

func partition(head *ListNode, x int) *ListNode {
	hook1 := &ListNode{Val: -1000, Next: head}
	hook2 := &ListNode{Val: 0, Next: nil}
	cur := hook1
	prev := hook1
	node2 := hook2
	for cur != nil {
		if cur.Val >= x {
			node2.Next = cur
			node2 = node2.Next
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	prev.Next = hook2.Next
	node2.Next = nil
	return hook1.Next
}

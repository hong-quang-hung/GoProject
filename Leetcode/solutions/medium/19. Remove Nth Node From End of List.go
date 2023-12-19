package medium

import "fmt"

// Reference: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func init() {
	Solutions[19] = func() {
		fmt.Println(`Input: head = [1,2,3,4,5], n = 2`)
		fmt.Println(`Output:`, removeNthFromEnd(S2ListNode(`[1,2,3,4,5]`), 2))
		fmt.Println(`Input: head = [1,2], n = 1`)
		fmt.Println(`Output:`, removeNthFromEnd(S2ListNode(`[1,2]`), 1))
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	for head != nil {
		clone := head
		nodes = append(nodes, clone)
		head = head.Next
	}

	nodes = append(nodes, nil)
	l := len(nodes)
	if n == l-1 {
		return nodes[1]
	}

	nodes[l-n-2].Next = nodes[l-n]
	return nodes[0]
}

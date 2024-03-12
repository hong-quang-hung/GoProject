package medium

import "fmt"

// Reference: https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
func init() {
	Solutions[1171] = func() {
		fmt.Println(`Input: head = [1,2,-3,3,1]`)
		fmt.Println(`Output:`, removeZeroSumSublists(S2ListNode("[1,2,-3,3,1]")))
		fmt.Println(`Input: head = [1,2,3,-3,4]`)
		fmt.Println(`Output:`, removeZeroSumSublists(S2ListNode("[1,2,3,-3,4]")))
		fmt.Println(`Input: head = [1,2,3,-3,-2]`)
		fmt.Println(`Output:`, removeZeroSumSublists(S2ListNode("[1,2,3,-3,-2]")))
	}
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	temp := &ListNode{Val: 0, Next: head}
	prefix := 0
	m := make(map[int]*ListNode)
	m[0] = temp
	for node := temp; node != nil; node = node.Next {
		prefix += node.Val
		m[prefix] = node
	}

	prefix = 0
	for node := temp; node != nil; node = node.Next {
		prefix += node.Val
		node.Next = m[prefix].Next
	}
	return temp.Next
}

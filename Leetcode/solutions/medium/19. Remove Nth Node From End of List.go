package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func Leetcode_Remove_Nth_From_End() {
	fmt.Println("Input: head = [1,2,3,4,5], n = 2")
	fmt.Println("Output:", removeNthFromEnd(utils.S2ListNode("[1,2,3,4,5]"), 2))
	fmt.Println("Input: head = [1,2], n = 1")
	fmt.Println("Output:", removeNthFromEnd(utils.S2ListNode("[1,2]"), 1))
	fmt.Println("Input: head = [1], n = 1")
	fmt.Println("Output:", removeNthFromEnd(utils.S2ListNode("[1]"), 1))
}

func removeNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	nodes := []*types.ListNode{}
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
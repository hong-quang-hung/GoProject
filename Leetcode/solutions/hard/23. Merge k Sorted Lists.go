package hard

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/add-two-numbers/
func Leetcode_Merge_K_Lists() {
	fmt.Println("Input: lists = [[1,4,5],[1,3,4],[2,6]]")
	fmt.Println("Output:", mergeKLists([]*types.ListNode{types.NewListNode(1, 4, 5), types.NewListNode(1, 3, 4), types.NewListNode(2, 6)}))
	fmt.Println("Input: lists = [[]]")
	fmt.Println("Output:", mergeKLists([]*types.ListNode{types.NewListNode()}))
}

func mergeKLists(lists []*types.ListNode) *types.ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	interval := 1
	for interval < length {
		for i := 0; i < length-interval; i += interval * 2 {
			lists[i] = mergeTwoNode(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
}

func mergeTwoNode(node1 *types.ListNode, node2 *types.ListNode) *types.ListNode {
	root := new(types.ListNode)
	temp := root
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			temp.Next = node2
			node2 = node1
			node1 = temp.Next.Next
		} else {
			temp.Next = node1
			node1 = node1.Next
		}
		temp = temp.Next
	}

	if node1 == nil {
		temp.Next = node2
	} else {
		temp.Next = node1
	}
	return root.Next
}

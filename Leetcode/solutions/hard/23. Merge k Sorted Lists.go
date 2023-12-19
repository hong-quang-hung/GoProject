package hard

import "fmt"

// Reference: https://leetcode.com/problems/merge-k-sorted-lists/
func init() {
	Solutions[23] = func() {
		fmt.Println(`Input: lists = [[1,4,5],[1,3,4],[2,6]]`)
		fmt.Println(`Output:`, mergeKLists([]*ListNode{S2ListNode(`[1,4,5]`), S2ListNode(`[1,3,4]`), S2ListNode(`[2,6]`)}))
		fmt.Println(`Input: lists = [[]]`)
		fmt.Println(`Output:`, mergeKLists([]*ListNode{S2ListNode(`[]`)}))
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
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

func mergeTwoNode(node1 *ListNode, node2 *ListNode) *ListNode {
	root := new(ListNode)
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

package medium

import "fmt"

// Reference: https://leetcode.com/problems/merge-in-between-linked-lists/
func init() {
	Solutions[1669] = func() {
		fmt.Println(`Input: list1 = [10,1,13,6,9,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]`)
		fmt.Println(`Output:`, mergeInBetween(S2ListNode("[10,1,13,6,9,5]"), 3, 4, S2ListNode("[1000000,1000001,1000002]")))
		fmt.Println(`Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]`)
		fmt.Println(`Output:`, mergeInBetween(S2ListNode("[0,1,2,3,4,5,6]"), 2, 5, S2ListNode("[1000000,1000001,1000002,1000003,1000004]")))
	}
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	p, q := list1, list1
	for ; a > 1; a-- {
		p = p.Next
	}

	for ; b > 0; b-- {
		q = q.Next
	}

	p.Next = list2
	for p.Next != nil {
		p = p.Next
	}

	p.Next = q.Next
	q.Next = nil
	return list1
}

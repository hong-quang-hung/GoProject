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
	return head
}

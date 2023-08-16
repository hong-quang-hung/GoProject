package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/intersection-of-two-linked-lists/
func init() {
	Solutions[160] = func() {
		fmt.Println("Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3")
		a := S2ListNode("[4,1]")
		b := S2ListNode("[5,6,1]")
		c := S2ListNode("[8,4,5]")
		a.Next = c
		b.Next = c
		fmt.Println("Output:", SListNode(getIntersectionNode(a, b)))
	}
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	flag := true
	for headA != nil && headB != nil {
		if headA != headB {
			if flag {
				headA = headA.Next
			} else {
				headB = headB.Next
			}
		} else {
			return headA
		}
	}
	return nil
}

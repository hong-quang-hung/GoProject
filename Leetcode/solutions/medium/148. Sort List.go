package medium

import "fmt"

// Reference: https://leetcode.com/problems/sort-list/
func init() {
	Solutions[148] = func() {
		fmt.Println("Input: head = [4,2,1,3]")
		fmt.Println("Output:", SListNode(sortList(S2ListNode("[4,2,1,3]"))))
		fmt.Println("Input: head = [-1,5,3,4,0]")
		fmt.Println("Output:", SListNode(sortList(S2ListNode("[-1,5,3,4,0]"))))
		fmt.Println("Input: head = [4,19,14,5,-3,1,8,5,11,15]")
		fmt.Println("Output:", SListNode(sortList(S2ListNode("[4,19,14,5,-3,1,8,5,11,15]"))))
	}
}

func sortList(head *ListNode) *ListNode {
	temp := head
	for temp != nil {
		next := temp
		for next.Next != nil && temp.Val >= next.Next.Val {
			next = next.Next
		}

		if temp != next && temp.Val != next.Val {
			temp.Val, next.Val = next.Val, temp.Val
		} else {
			if temp.Val < head.Val {
				temp = head
			} else {
				temp = temp.Next
			}
		}
	}
	return head
}

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
	if head == nil || head.Next == nil {
		return head
	}

	other := sortListDivide(head)
	head1 := sortList(head)
	head2 := sortList(other)
	return sortListMerge(head1, head2)
}

func sortListDivide(head *ListNode) *ListNode {
	slow := head
	fast := slow.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	newHead := slow.Next
	slow.Next = nil
	return newHead
}

func sortListMerge(p, q *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	for p != nil && q != nil {
		if p.Val < q.Val {
			temp.Next = p
			p = p.Next
		} else {
			temp.Next = q
			q = q.Next
		}
		temp = temp.Next
	}

	if p != nil {
		temp.Next = p
	} else {
		temp.Next = q
	}
	return res.Next
}

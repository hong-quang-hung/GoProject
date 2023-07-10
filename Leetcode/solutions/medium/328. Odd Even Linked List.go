package medium

import "fmt"

func init() {
	Solutions[328] = Leetcode_Odd_Even_List
}

// Reference: https://leetcode.com/problems/odd-even-linked-list/
func Leetcode_Odd_Even_List() {
	fmt.Println("Input: head = [1,2,3,4,5]")
	fmt.Println("Output:", SListNode(oddEvenList(S2ListNode("[1,2,3,4,5]"))))
	fmt.Println("Input: head = [2,1,3,5,6,4,7]")
	fmt.Println("Output:", SListNode(oddEvenList(S2ListNode("[2,1,3,5,6,4,7]"))))
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd, even := head, head.Next
	beforeEven := even
	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}

	odd.Next = beforeEven
	return head
}

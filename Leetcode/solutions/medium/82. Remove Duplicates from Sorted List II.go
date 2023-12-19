package medium

import "fmt"

// Reference: https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
func init() {
	Solutions[82] = func() {
		fmt.Println(`Input: head = [1,2,3,3,4,4,5]`)
		fmt.Println(`Output:`, SListNode(deleteDuplicates(S2ListNode(`[1,2,3,3,4,4,5]`))))
		fmt.Println(`Input: head = [1,1,1,2,3]`)
		fmt.Println(`Output:`, SListNode(deleteDuplicates(S2ListNode(`[1,1,1,2,3]`))))
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	temp := head
	for temp.Next != nil && temp.Val == temp.Next.Val {
		temp = temp.Next
	}

	if head != temp && head.Val == temp.Val {
		head = deleteDuplicates(temp.Next)
	} else {
		head = temp
		head.Next = deleteDuplicates(head.Next)
	}
	return head
}

package easy

import "fmt"

// Reference: https://leetcode.com/problems/palindrome-linked-list/
func init() {
	Solutions[234] = func() {
		fmt.Println("Input: head = [1,2,2,1]")
		fmt.Println("Output:", isPalindromeII(S2ListNode("[1,2,2,1]")))
		fmt.Println("Input: head = [1,2]")
		fmt.Println("Output:", isPalindromeII(S2ListNode("[1,2]")))
		fmt.Println("Input: head = [1,2,3,2,1]")
		fmt.Println("Output:", isPalindromeII(S2ListNode("[1,2,3,2,1]")))
	}
}

func isPalindromeII(head *ListNode) bool {
	reverse, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		next := reverse.Next
		reverse.Next = nil
		next.Next = reverse
		reverse = next
	}

	fmt.Println(SListNode(reverse), SListNode(slow), SListNode(fast))
	return false
}

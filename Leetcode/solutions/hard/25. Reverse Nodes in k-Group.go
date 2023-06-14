package hard

import "fmt"

// Reference: https://leetcode.com/problems/reverse-nodes-in-k-group/
func Leetcode_Reverse_K_Group() {
	fmt.Println("Input: head = [1,2,3,4,5], k = 2")
	fmt.Println("Output:", reverseKGroup(S2ListNode("[1,2,3,4,5]"), 2))
	fmt.Println("Input: head = [1,2,3,4,5], k = 3")
	fmt.Println("Output:", reverseKGroup(S2ListNode("[1,2,3,4,5]"), 3))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	temp, count := new(ListNode), 0
	temp.Next = head
	curr, next, prev := temp, temp, temp
	for curr.Next != nil {
		count++
		curr = curr.Next
	}

	for count >= k {
		curr = prev.Next
		next = curr.Next
		for i := 1; i < k; i++ {
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
			next = curr.Next
		}

		prev = curr
		count -= k
	}
	return temp.Next
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/linked-list-cycle-ii/
func init() {
	Solutions[142] = func() {
		fmt.Println("Input: head = [3,2,0,-4], pos = 1")
		node := S2ListNode("[3,2,0,-4]")
		pos := node.Next
		node.Next.Next.Next = pos
		fmt.Println("Output:", detectCycle(node))
	}
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

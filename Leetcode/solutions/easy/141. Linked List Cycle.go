package easy

import "fmt"

func init() {
	Solutions[141] = Leetcode_Has_Cycle
}

// Reference: https://leetcode.com/problems/linked-list-cycle/
func Leetcode_Has_Cycle() {
	fmt.Println("Input: head = [3,2,0,-4], pos = 1")
	node := S2ListNode("[3,2,0,-4]")
	post := node.Next
	last := node.Last()
	last.Next = post
	fmt.Println("Output:", hasCycle(node))
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
	}
	return false
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
func Leetcode_Pair_Sum() {
	fmt.Println("Input: head = [5,4,2,1]")
	fmt.Println("Output:", pairSum(S2ListNode("[5,4,2,1]")))
	fmt.Println("Input: head = [4,2,2,3]")
	fmt.Println("Output:", pairSum(S2ListNode("[4,2,2,3]")))
	fmt.Println("Input: head = [1,100000]")
	fmt.Println("Output:", pairSum(S2ListNode("[1,100000]")))
}

func pairSum(head *ListNode) int {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var prev *ListNode
	for slow != nil {
		nextNode := slow.Next
		slow.Next = prev
		prev = slow
		slow = nextNode
	}

	res, start := 0, head
	for prev != nil {
		res = max(res, prev.Val+start.Val)
		prev = prev.Next
		start = start.Next
	}
	return res
}

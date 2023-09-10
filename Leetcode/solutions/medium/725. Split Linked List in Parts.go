package medium

import "fmt"

// Reference: https://leetcode.com/problems/split-linked-list-in-parts/
func init() {
	Solutions[725] = func() {
		fmt.Println("Input: head = [1,2,3], k = 5")
		fmt.Println("Output:", splitListToParts(S2ListNode("[1,2,3]"), 5))
		fmt.Println("Input: head = [1,2,3,4,5,6,7,8,9,10], k = 3")
		fmt.Println("Output:", splitListToParts(S2ListNode("[1,2,3,4,5,6,7,8,9,10]"), 3))
	}
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	average, remainder := length/k, length%k
	var prev *ListNode
	for node, i := head, 0; node != nil && i < k; i++ {
		res[i] = node
		count := average
		if remainder > 0 {
			count++
		}
		for j := 0; j < count; j++ {
			prev = node
			node = node.Next
		}

		prev.Next = nil
		remainder--
	}
	return res
}

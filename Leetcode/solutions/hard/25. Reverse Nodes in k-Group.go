package hard

import "fmt"

// Reference: https://leetcode.com/problems/maximum-subsequence-score/
func Leetcode_Reverse_K_Group() {
	fmt.Println("Input: head = [1,2,3,4,5], k = 2")
	fmt.Println("Output:", reverseKGroup(S2ListNode("[1,2,3,4,5]"), 2))
	fmt.Println("Input: head = [1,2,3,4,5], k = 3")
	fmt.Println("Output:", reverseKGroup(S2ListNode("[1,2,3,4,5]"), 3))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	temp, count := head, 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	return head
}

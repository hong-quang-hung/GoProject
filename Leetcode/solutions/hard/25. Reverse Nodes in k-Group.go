package hard

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/maximum-subsequence-score/
func Leetcode_Reverse_K_Group() {
	fmt.Println("Input: head = [1,2,3,4,5], k = 2")
	fmt.Println("Output:", reverseKGroup(utils.S2ListNode("[1,2,3,4,5]"), 2))
	fmt.Println("Input: head = [1,2,3,4,5], k = 3")
	fmt.Println("Output:", reverseKGroup(utils.S2ListNode("[1,2,3,4,5]"), 3))
}

func reverseKGroup(head *types.ListNode, k int) *types.ListNode {
	temp, count := head, 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	return head
}
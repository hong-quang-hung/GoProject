package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/swap-nodes-in-pairs/
func Leetcode_Swap_Pairs() {
	fmt.Println("Input: head = [1,2,3,4]")
	fmt.Println("Output:", swapPairs(utils.S2ListNode("[1,2,3,4,5]")))
	fmt.Println("Input: head = [1]")
	fmt.Println("Output:", swapPairs(utils.S2ListNode("[1]")))
}

func swapPairs(head *types.ListNode) *types.ListNode {
	if head != nil && head.Next != nil {
		head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head
	}
	return head
}

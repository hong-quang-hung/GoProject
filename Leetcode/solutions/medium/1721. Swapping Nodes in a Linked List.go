package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
func Leetcode_Swap_Nodes() {
	fmt.Println("Input: head = [1,2,3,4,5], k = 2")
	fmt.Println("Output:", swapNodes(utils.S2ListNode("[1,2,3,4,5]"), 2))
	fmt.Println("Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5")
	fmt.Println("Output:", swapNodes(utils.S2ListNode("[7,9,6,6,7,8,3,0,9,5]"), 5))
}

func swapNodes(head *types.ListNode, k int) *types.ListNode {
	temp1, i, n := head, 0, 0

	var left *types.ListNode
	for temp1 != nil {
		i++
		n++
		if i == k {
			left = temp1
		}
		temp1 = temp1.Next
	}

	i = 0
	var right *types.ListNode
	temp2 := head
	for temp2 != nil {
		i++
		if i == n-k+1 {
			right = temp2
		}
		temp2 = temp2.Next
	}

	temVal := left.Val
	left.Val = right.Val
	right.Val = temVal
	return head
}

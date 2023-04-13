package medium

import (
	"fmt"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/add-two-numbers/
func Leetcode_Add_Two_Numbers() {
	fmt.Println("Input: l1 = [2,4,3], l2 = [5,6,4]")
	fmt.Println("Output:", addTwoNumbers(types.NewListNode(2, 4, 3), types.NewListNode(5, 6, 4)))
	fmt.Println("Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]")
	fmt.Println("Output:", addTwoNumbers(types.NewListNode(9, 9, 9, 9, 9, 9, 9), types.NewListNode(9, 9, 9, 9)))
}

func addTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	p1, p2, carry := l1, l2, 0
	for p1 != nil || p2 != nil || carry == 1 {
		if p2 == nil {
			p1.Val = p1.Val + carry
		} else {
			p1.Val = p1.Val + p2.Val + carry
		}

		if p1.Val >= 10 {
			p1.Val -= 10
			carry = 1
		} else {
			carry = 0
		}

		if p2 != nil {
			p2 = p2.Next
		}
		if p1.Next == nil && (carry == 1 || p2 != nil) {
			p1.Next = new(types.ListNode)
		}
		p1 = p1.Next
	}
	return l1
}

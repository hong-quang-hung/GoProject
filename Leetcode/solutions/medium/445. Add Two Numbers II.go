package medium

import "fmt"

// Reference: https://leetcode.com/problems/add-two-numbers-ii/
func init() {
	Solutions[445] = func() {
		fmt.Println("Input: l1 = [7,2,4,3], l2 = [5,6,4]")
		fmt.Println("Output:", addTwoNumbersII(S2ListNode("[7,2,4,3]"), S2ListNode("[5,6,4]")))
		fmt.Println("Input: l1 = [2,4,3], l2 = [5,6,4]")
		fmt.Println("Output:", addTwoNumbersII(S2ListNode("[2,4,3]"), S2ListNode("[5,6,4]")))
	}
}

func addTwoNumbersII(first *ListNode, second *ListNode) *ListNode {
	first = addTwoNumbersIIReverse(first)
	second = addTwoNumbersIIReverse(second)
	carry := 0

	var third, last *ListNode = nil, nil
	for first != nil || second != nil || carry != 0 {
		sum := carry
		if first != nil {
			sum += first.Val
		}

		if second != nil {
			sum += second.Val
		}

		temp := &ListNode{Val: sum % 10, Next: nil}
		if third == nil {
			third = temp
			last = temp
		} else {
			last.Next = temp
			last = last.Next
		}

		carry = sum / 10
		if first != nil {
			first = first.Next
		}

		if second != nil {
			second = second.Next
		}
	}
	return addTwoNumbersIIReverse(third)
}

func addTwoNumbersIIReverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := addTwoNumbersIIReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

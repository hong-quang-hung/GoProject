package medium

import "fmt"

// Reference: https://leetcode.com/problems/copy-list-with-random-pointer/
func init() {
	Solutions[138] = func() {
		fmt.Println("Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]")
		node0, node1, node2, node3, node4 := new(RandomNode), new(RandomNode), new(RandomNode), new(RandomNode), new(RandomNode)
		node0.Val, node1.Val, node2.Val, node3.Val, node4.Val = 7, 13, 11, 10, 1
		node0.Next, node1.Next, node2.Next, node3.Next = node1, node2, node3, node4
		node1.Random, node2.Random, node3.Random, node4.Random = node0, node4, node2, node0
		fmt.Println("Output:", copyRandomList(node0))
	}
}

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}

	temp1 := head
	for temp1 != nil {
		newNode := &RandomNode{
			Val:    temp1.Val,
			Next:   temp1.Next,
			Random: temp1.Random,
		}

		temp1.Next = newNode
		temp1 = newNode.Next
	}

	temp1 = head
	for temp1 != nil {
		temp2 := temp1.Next
		if temp1.Random == nil {
			temp2.Random = nil
		} else {
			temp2.Random = temp1.Random.Next
		}
		temp1 = temp1.Next.Next
	}

	temp1 = head
	res := head.Next
	for temp1 != nil {
		temp2 := temp1.Next
		temp1.Next = temp2.Next
		if temp1.Next != nil {
			temp2.Next = temp1.Next.Next
		} else {
			temp2.Next = nil
		}
		temp1 = temp1.Next
	}
	return res
}

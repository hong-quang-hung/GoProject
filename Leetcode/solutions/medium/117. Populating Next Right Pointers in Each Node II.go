package medium

import "fmt"

// Reference: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
func init() {
	Solutions[117] = func() {
		fmt.Println(`Input: root = [1,2,3,4,5,null,7]`)
		root := new(NodeII)
		root.Val = 1
		root.Left = new(NodeII)
		root.Left.Val = 2
		root.Left.Left = new(NodeII)
		root.Left.Left.Val = 4
		root.Left.Right = new(NodeII)
		root.Left.Right.Val = 5
		root.Right = new(NodeII)
		root.Right.Val = 3
		root.Right.Left = nil
		root.Right.Right = new(NodeII)
		root.Right.Right.Val = 7
		fmt.Println(`Output:`, connect(root))
	}
}

type NodeII struct {
	Val   int
	Left  *NodeII
	Right *NodeII
	Next  *NodeII
}

func connect(root *NodeII) *NodeII {
	if root == nil {
		return nil
	}

	q := []*NodeII{root}
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if size > 1 {
				cur.Next = q[0]
			}
			size--
		}
	}
	return root
}

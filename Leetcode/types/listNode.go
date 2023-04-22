package types

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val ...int) *ListNode {
	if len(val) == 0 {
		return nil
	}
	if len(val) == 1 {
		return &ListNode{Val: val[0], Next: nil}
	}
	node := &ListNode{Val: val[0], Next: nil}
	node.Next = NewListNode(val[1:]...)
	return node
}

// ListNode must not cycle
func (t *ListNode) Last() *ListNode {
	last := t
	for last.Next != nil {
		last = last.Next
	}
	return last
}

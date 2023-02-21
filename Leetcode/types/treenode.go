package types

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeNodeBase(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

package types

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LazyNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func LazyNodeValue(val int, valf int, valr int) *TreeNode {
	return &TreeNode{Val: val, Left: LazyNode(valf), Right: LazyNode(valr)}
}

func LazyNodeAll(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func (root TreeNode) Println() {
	fmt.Println("[]")
}

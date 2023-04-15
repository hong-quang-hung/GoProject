package types

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nilStr string = "null"

func LazyNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func LazyNodeValue(val int, valf int, valr int) *TreeNode {
	return &TreeNode{Val: val, Left: LazyNode(valf), Right: LazyNode(valr)}
}

func LazyNodeAll(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

var maxDepth int
var strDepth []string

func (root *TreeNode) Sprint() string {
	strDepth = strDepth[:cap(strDepth)]
	maxDepth = 0
	getNode(root, 0)
	strDepth = strDepth[0 : maxDepth+1]
	return "[" + strings.Join(strDepth, ",") + "]"
}

func getNode(node *TreeNode, depth int) {
	var val string
	if node == nil {
		val = nilStr
	} else {
		val = strconv.Itoa(node.Val)
	}

	if len(strDepth) > depth {
		strDepth[depth] = strDepth[depth] + "," + val
	} else {
		strDepth = append(strDepth, val)
	}

	if node == nil {
		return
	}

	maxDepth = depth
	getNode(node.Left, depth+1)
	getNode(node.Right, depth+1)
}

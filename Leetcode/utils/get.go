package utils

import (
	"strconv"
	"strings"

	"leetcode.com/Leetcode/types"
)

func Iff[T any](c bool, a T, b T) T {
	if c {
		return a
	}
	return b
}

func IffNil[T any](c any, a T, b T) T {
	if c == nil {
		return a
	}
	return b
}

// Convert input string to ListNode
func S2ListNode(s string) *types.ListNode {
	return types.NewListNode(S2SliceInt(s)...)
}

// Returning the TreeNode string
// do not use when ListNode is cycle...
func SListNode(head *types.ListNode) string {
	if head == nil {
		return "[]"
	}

	var s strings.Builder
	s.WriteString("[")
	temp := head
	for {
		s.WriteString(strconv.Itoa(temp.Val))
		temp = temp.Next
		if temp == nil {
			break
		} else {
			s.WriteString(",")
		}
	}
	s.WriteString("]")
	return s.String()
}

// Convert input string to TreeNode
func S2TreeNode(s string) *types.TreeNode {
	values := S2SliceString(s)
	n, queue := len(values), []*types.TreeNode{}
	if n == 0 {
		return nil
	}

	flag := true
	rootVal, _ := strconv.ParseInt(values[0], 10, 0)
	root := types.LazyNode(int(rootVal))
	queue = append(queue, []*types.TreeNode{root, root}...)
	for i := 1; i < n; i++ {
		if values[i] != nilStr {
			val, _ := strconv.ParseInt(values[i], 10, 0)
			if flag {
				queue[0].Left = types.LazyNode(int(val))
				queue = append(queue, []*types.TreeNode{queue[0].Left, queue[0].Left}...)
			} else {
				queue[0].Right = types.LazyNode(int(val))
				queue = append(queue, []*types.TreeNode{queue[0].Right, queue[0].Right}...)
			}
		}
		flag = !flag
		queue = queue[1:]
	}
	return root
}

// Returning the TreeNode string
func STreeNode(root *types.TreeNode) string {
	strDepth, maxDepth := []string{}, 0
	printTreeNode(root, &strDepth, &maxDepth, 0)
	strDepth = strDepth[0:(maxDepth + 1)]
	return "[" + strings.Join(strDepth, ",") + "]"
}

func printTreeNode(node *types.TreeNode, strDepth *[]string, maxDepth *int, depth int) {
	var val string
	if node == nil {
		val = nilStr
	} else {
		val = strconv.Itoa(node.Val)
	}

	if len(*strDepth) > depth {
		(*strDepth)[depth] = (*strDepth)[depth] + "," + val
	} else {
		*strDepth = append(*strDepth, val)
	}

	if node == nil {
		return
	}

	*maxDepth = Max(*maxDepth, depth)
	printTreeNode(node.Left, strDepth, maxDepth, depth+1)
	printTreeNode(node.Right, strDepth, maxDepth, depth+1)
}

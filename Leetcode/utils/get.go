package utils

import (
	"fmt"
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

// Print Any array
func SAny[L Any](a []*L) string {
	if len(a) == 0 {
		return nilSlice
	}

	var s strings.Builder
	s.WriteString(sliceStart)
	s.WriteString(printAny(a[0]))
	for _, o := range a[1:] {
		s.WriteString(commaSpaceString)
		s.WriteString(printAny(o))
	}
	s.WriteString(sliceEnd)
	return s.String()
}

// Convert input string to ListNode
func S2ListNode(s string) *types.ListNode {
	return types.NewListNode(S2SliceInt(s)...)
}

// Returning the TreeNode string
// do not use when ListNode is cycle...
func SListNode(head *types.ListNode) string {
	if head == nil {
		return nilSlice
	}

	var s strings.Builder
	s.WriteString(sliceStart)
	temp := head
	for {
		s.WriteString(strconv.Itoa(temp.Val))
		temp = temp.Next
		if temp == nil {
			break
		} else {
			s.WriteString(commaSpaceString)
		}
	}
	s.WriteString(sliceEnd)
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
	if root == nil {
		return nilSlice
	}

	strDepth, maxDepth := [][]string{}, 0
	printTreeNode(root, &strDepth, &maxDepth, 0)
	strDepth = strDepth[0:(maxDepth + 1)]

	temp := strDepth[maxDepth]
	idx := len(temp) - 1
	for idx >= 0 && temp[idx] == nilStr {
		idx--
	}
	strDepth[maxDepth] = temp[:idx+1]

	str := new(strings.Builder)
	str.WriteString(sliceStart)
	str.WriteString(strings.Join(strDepth[0], commaSpaceString))
	for _, sd := range strDepth[1:] {
		str.WriteString(commaSpaceString)
		str.WriteString(strings.Join(sd, commaSpaceString))
	}
	str.WriteString(sliceEnd)
	return str.String()
}

func printTreeNode(node *types.TreeNode, strDepth *[][]string, maxDepth *int, depth int) {
	var val string
	if node == nil {
		val = nilStr
	} else {
		val = strconv.Itoa(node.Val)
	}

	if len(*strDepth) > depth {
		(*strDepth)[depth] = append((*strDepth)[depth], val)
	} else {
		*strDepth = append(*strDepth, []string{val})
	}

	if node == nil {
		return
	}

	*maxDepth = Max(*maxDepth, depth)
	printTreeNode(node.Left, strDepth, maxDepth, depth+1)
	printTreeNode(node.Right, strDepth, maxDepth, depth+1)
}

// Implement Print
func printAny(a any) string {
	switch t := a.(type) {
	case *types.TreeNode:
		return STreeNode(t)
	case *types.ListNode:
		return SListNode(t)
	default:
		return fmt.Sprint(t)
	}
}

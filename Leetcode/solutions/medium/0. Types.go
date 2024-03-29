package medium

import (
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

const (
	mod = 1e9 + 7
)

type (
	Any       = utils.Any
	Node      = types.Node
	ListNode  = types.ListNode
	TreeNode  = types.TreeNode
	TrieNode  = types.TrieNode
	TrieNodes = types.TrieNodes
	UnionFind = types.UnionFind
	DSU       = types.DSU
	MinHeap   = types.MinHeap
	MaxHeap   = types.MaxHeap
	Solution  = types.Solution
)

var (
	Solutions       Solution = make(Solution)
	KnightDirection          = [][2]int{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}
	BoardDirection           = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func SAny[L Any](a []*L) string {
	return utils.SAny(a)
}

func S2ListNode(s string) *ListNode {
	return utils.S2ListNode(s)
}

func SListNode(head *ListNode) string {
	return utils.SListNode(head)
}

func STreeNode(root *TreeNode) string {
	return utils.STreeNode(root)
}

func S2TreeNode(s string) *TreeNode {
	return utils.S2TreeNode(s)
}

func S2SliceInt(s string) []int {
	return utils.S2SliceInt(s)
}

func S2SoSliceInt(s string) [][]int {
	return utils.S2SoSliceInt(s)
}

func S2SoSliceByte(s string) [][]byte {
	return utils.S2SoSliceByte(s)
}

func NewUnionFind(n int) *UnionFind {
	return types.NewUnionFind(n)
}

func NewDSU(n int) *DSU {
	return types.NewDSU(n)
}

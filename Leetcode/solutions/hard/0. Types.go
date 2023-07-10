package hard

import (
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

const mod = 1_000_000_007

type (
	Node      = types.Node
	ListNode  = types.ListNode
	TreeNode  = types.TreeNode
	TrieNode  = types.TrieNode
	UnionFind = types.UnionFind
	MinHeap   = types.MinHeap
	MaxHeap   = types.MaxHeap
	Solution  = types.Solution
)

var Solutions Solution = make(Solution)

func S2ListNode(s string) *ListNode {
	return utils.S2ListNode(s)
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

func NewUnionFind(size int) *UnionFind {
	return types.NewUnionFind(size)
}

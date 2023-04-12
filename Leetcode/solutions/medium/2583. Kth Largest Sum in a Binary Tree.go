package medium

import (
	"fmt"
	"sort"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/
func Leetcode_Kth_Largest_Level_Sum() {
	fmt.Println("Input: root = [5,8,9,2,1,3,7,4,6], k = 2")
	fmt.Println("Output:", kthLargestLevelSum(types.LazyNodeAll(5, types.LazyNodeAll(8, types.LazyNodeValue(2, 4, 6), types.LazyNode(1)), types.LazyNodeValue(9, 3, 7)), 2))
}

func kthLargestLevelSum(root *types.TreeNode, k int) int64 {
	sumk := make([]int64, 0)
	sumNodeRecursive(root, &sumk, 0)
	if len(sumk) < k {
		return -1
	}

	sort.Slice(sumk, func(i, j int) bool { return sumk[i] > sumk[j] })
	return sumk[k-1]
}

func sumNodeRecursive(root *types.TreeNode, sumK *[]int64, k int) {
	if root == nil {
		return
	}

	if len(*sumK) <= k {
		*sumK = append(*sumK, 0)
	}
	(*sumK)[k] += int64(root.Val)

	if root.Left != nil {
		sumNodeRecursive(root.Left, sumK, k+1)
	}
	if root.Right != nil {
		sumNodeRecursive(root.Right, sumK, k+1)
	}
}

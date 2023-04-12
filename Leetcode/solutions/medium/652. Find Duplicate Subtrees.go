package medium

import (
	"fmt"
	"strconv"

	"leetcode.com/Leetcode/types"
)

// Reference: https://leetcode.com/problems/find-duplicate-subtrees/
func Leetcode_Find_Duplicate_Subtrees() {
	fmt.Println("Input: root = [1,2,3,4,null,2,4,null,null,4]")
	fmt.Println("Output:", findDuplicateSubtrees(types.LazyNodeAll(1, types.LazyNodeAll(2, types.LazyNode(4), nil), types.LazyNodeAll(3, types.LazyNodeAll(2, types.LazyNode(4), nil), types.LazyNode(4)))))
}

func findDuplicateSubtrees(root *types.TreeNode) []*types.TreeNode {
	res := make([]*types.TreeNode, 0)
	tripletToID := make(map[string]int)
	cnt := make(map[int]int)
	findDuplicateSubtreesTraversal(root, tripletToID, cnt, &res)
	return res
}

func findDuplicateSubtreesTraversal(node *types.TreeNode, tripletToID map[string]int, cnt map[int]int, res *[]*types.TreeNode) int {
	if node == nil {
		return 0
	}

	triplet := strconv.Itoa(findDuplicateSubtreesTraversal(node.Left, tripletToID, cnt, res)) + "," + strconv.Itoa(node.Val) + "," + strconv.Itoa(findDuplicateSubtreesTraversal(node.Right, tripletToID, cnt, res))
	if _, c := tripletToID[triplet]; !c {
		tripletToID[triplet] = len(tripletToID) + 1
	}

	id := tripletToID[triplet]
	cnt[id]++
	if cnt[id] == 2 {
		*res = append(*res, node)
	}
	return id
}

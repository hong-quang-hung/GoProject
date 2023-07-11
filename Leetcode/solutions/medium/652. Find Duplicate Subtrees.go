package medium

import (
	"fmt"
	"strconv"
)

// Reference: https://leetcode.com/problems/find-duplicate-subtrees/
func init() {
	Solutions[652] = func() {
		fmt.Println("Input: root = [1,2,3,4,null,2,4,null,null,4]")
		fmt.Println("Output:", findDuplicateSubtrees(S2TreeNode("[1,2,3,4,null,2,4,null,null,4]")))
	}
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	tripletToID := make(map[string]int)
	cnt := make(map[int]int)
	findDuplicateSubtreesTraversal(root, tripletToID, cnt, &res)
	return res
}

func findDuplicateSubtreesTraversal(node *TreeNode, tripletToID map[string]int, cnt map[int]int, res *[]*TreeNode) int {
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

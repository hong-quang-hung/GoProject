package easy

import (
	"fmt"
	"strconv"
	"strings"
)

// Reference: https://leetcode.com/problems/construct-string-from-binary-tree/
func init() {
	Solutions[606] = func() {
		fmt.Println("Input: root = [1,2,3,4]")
		fmt.Println("Output:", tree2str(S2TreeNode("[1,2,3,4]")))
		fmt.Println("Input: root = [1,2,3,null,4]")
		fmt.Println("Output:", tree2str(S2TreeNode("[1,2,3,null,4]")))
	}
}

func tree2str(root *TreeNode) string {
	res := strings.Builder{}
	res.WriteString(strconv.Itoa(root.Val))
	isLeft, isRight := root.Left != nil, root.Right != nil
	if isLeft {
		res.WriteString("(")
		res.WriteString(tree2str(root.Left))
		res.WriteString(")")
	} else if isRight {
		res.WriteString("()")
	}

	if isRight {
		res.WriteString("(")
		res.WriteString(tree2str(root.Right))
		res.WriteString(")")
	}
	return res.String()
}

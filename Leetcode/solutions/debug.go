package solutions

import (
	"fmt"

	T "leetcode.com/Leetcode/types"
)

func Leetcode_Two_Sum() {
	target := 9
	nums := [...]int{2, 7, 11, 15}
	fmt.Println("Input: nums = [2,7,11,15], target = 9")
	fmt.Println("Output:", twoSum(nums[:], target))
}

func LeetCode_Max_Depth() {
	var root *T.TreeNode

	root = &T.TreeNode{Val: 3, Left: T.TreeNodeBase(9), Right: &T.TreeNode{Val: 20, Left: T.TreeNodeBase(15), Right: T.TreeNodeBase(7)}}
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output: ", maxDepth(root))

	root = &T.TreeNode{Val: 1, Left: nil, Right: T.TreeNodeBase(9)}
	fmt.Println("Input: root = [1,null,2]")
	fmt.Println("Output: ", maxDepth(root))
}

package medium

import "fmt"

// Reference: https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
func init() {
	Solutions[863] = func() {
		root := S2TreeNode("[3,5,1,6,2,0,8,null,null,7,4]")
		target := root.Left
		fmt.Println("Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2")
		fmt.Println("Output:", distanceK(root, target, 2))

		root = S2TreeNode("[1]")
		target = root
		fmt.Println("Input: root = [1], target = 1, k = 3")
		fmt.Println("Output:", distanceK(root, target, 3))
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	res := []int{}
	distanceKSolved(root, target, k, &res)
	return res
}

func distanceKSolved(root *TreeNode, target *TreeNode, k int, result *[]int) int {
	if root == nil {
		return -1
	}

	if root == target {
		distanceKFind(root, 0, k, result)
		return 1
	}

	leftDistance := distanceKSolved(root.Left, target, k, result)
	rightDistance := distanceKSolved(root.Right, target, k, result)
	if leftDistance != -1 {
		if leftDistance == k {
			*result = append(*result, root.Val)
		} else {
			distanceKFind(root.Right, leftDistance+1, k, result)
		}
		return leftDistance + 1
	}

	if rightDistance != -1 {
		if rightDistance == k {
			*result = append(*result, root.Val)
		} else {
			distanceKFind(root.Left, rightDistance+1, k, result)
		}
		return rightDistance + 1
	}
	return -1
}

func distanceKFind(root *TreeNode, distance, k int, result *[]int) {
	if root == nil {
		return
	}

	if distance > k {
		return
	}

	if distance == k {
		*result = append(*result, root.Val)
		return
	}

	distanceKFind(root.Left, distance+1, k, result)
	distanceKFind(root.Right, distance+1, k, result)
}

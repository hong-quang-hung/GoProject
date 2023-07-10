package medium

import (
	"fmt"
	"math"
)

func init() {
	Solutions[1161] = Leetcode_Max_LevelSum
}

// Reference: https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
func Leetcode_Max_LevelSum() {
	fmt.Println("Input: root = [1,7,0,7,-8,null,null]")
	fmt.Println("Output:", maxLevelSum(S2TreeNode("[1,7,0,7,-8,null,null]")))
	fmt.Println("Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]")
	fmt.Println("Output:", maxLevelSum(S2TreeNode("[989,null,10250,98693,-89388,null,null,null,-32127]")))
}

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	res, level, maxSum := 0, 0, math.MinInt
	for len(queue) > 0 {
		level++
		sum, w := 0, len(queue)
		for i := 0; i < w; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[w:]
		if maxSum < sum {
			res = level
			maxSum = sum
		}
	}
	return res
}

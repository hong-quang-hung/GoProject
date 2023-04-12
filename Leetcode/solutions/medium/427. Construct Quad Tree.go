package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/construct-quad-tree/
func Leetcode_Construct() {
	fmt.Println("Input: grid = [[0,1],[1,0]]")
	fmt.Println("Output:", construct(utils.S2SoSliceInt("[[0,1],[1,0]]")))
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return solveConstruct(grid, 0, 0, len(grid))
}

func solveConstruct(grid [][]int, x1, y1, length int) *Node {
	if length == 1 {
		return &Node{Val: grid[x1][y1] == 1, IsLeaf: true}
	}

	topLeft := solveConstruct(grid, x1, y1, length/2)
	topRight := solveConstruct(grid, x1, y1+length/2, length/2)
	bottomLeft := solveConstruct(grid, x1+length/2, y1, length/2)
	bottomRight := solveConstruct(grid, x1+length/2, y1+length/2, length/2)
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf && topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &Node{Val: topLeft.Val, IsLeaf: true}
	}

	return &Node{Val: false, IsLeaf: false, TopLeft: topLeft, TopRight: topRight, BottomLeft: bottomLeft, BottomRight: bottomRight}
}

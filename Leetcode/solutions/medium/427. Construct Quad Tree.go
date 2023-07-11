package medium

import "fmt"

// Reference: https://leetcode.com/problems/construct-quad-tree/
func init() {
	Solutions[427] = func() {
		fmt.Println("Input: grid = [[0,1],[1,0]]")
		fmt.Println("Output:", construct(S2SoSliceInt("[[0,1],[1,0]]")))
	}
}

type Cell struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Cell
	TopRight    *Cell
	BottomLeft  *Cell
	BottomRight *Cell
}

func construct(grid [][]int) *Cell {
	return solveConstruct(grid, 0, 0, len(grid))
}

func solveConstruct(grid [][]int, x1, y1, length int) *Cell {
	if length == 1 {
		return &Cell{Val: grid[x1][y1] == 1, IsLeaf: true}
	}

	topLeft := solveConstruct(grid, x1, y1, length/2)
	topRight := solveConstruct(grid, x1, y1+length/2, length/2)
	bottomLeft := solveConstruct(grid, x1+length/2, y1, length/2)
	bottomRight := solveConstruct(grid, x1+length/2, y1+length/2, length/2)
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf && topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &Cell{Val: topLeft.Val, IsLeaf: true}
	}

	return &Cell{Val: false, IsLeaf: false, TopLeft: topLeft, TopRight: topRight, BottomLeft: bottomLeft, BottomRight: bottomRight}
}

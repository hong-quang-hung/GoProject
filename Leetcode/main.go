package main

import (
	"fmt"

	"leetcode.com/Leetcode/solutions"
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

func main() {
	fmt.Println("..............................................................................................................................................")
	fmt.Println("Golang Leetcode...")
	solutions.Leetcode_Min_Path_Sum()
	fmt.Println("..............................................................................................................................................")
	GetRandomProblem()
	fmt.Println("..............................................................................................................................................")
	// TestPattern()
}

func GetRandomProblem() {
	Leetcode := &types.Leetcode{}
	Leetcode.SetTotal(2603)
	Leetcode.SetSolved(1, 2, 3, 4, 5, 9, 13, 14, 15, 20, 21, 23, 26, 27, 28, 35, 58, 67, 72, 101, 103, 104, 106, 109, 112, 121, 129, 142, 208, 211, 226, 290, 382, 427, 443, 502, 540, 567, 605)
	Leetcode.SetSolved(652, 783, 875, 912, 953, 958, 989, 1011, 1071, 1129, 1137, 1162, 1207, 1319, 1345, 1466, 1470, 1472, 1523, 1539, 1675, 2187, 2316, 2348, 2360, 2444, 2477, 2492, 2566)
	Leetcode.SetSolved(2567, 2570, 2571, 2574, 2575, 2576, 2578, 2579, 2580, 2582, 2583, 2586, 2587, 2588, 2591, 2592, 2595, 2596, 2597, 2600, 2601, 2602, 64)

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problems has been solved in Leetcode.")
	fmt.Println("Today, Number of Leetcode Problem is:", Leetcode.GetRandom())
}

func TestPattern() {
	fmt.Println("Input: nums = [0,0,0,2,0,0]")
	fmt.Println("Output:", utils.Sslice([]int{0, 0, 0, 2, 0, 0}))
	fmt.Println("Input: grid = \"[[1,0,0],[0,0,0],[0,0,0]]\"")
	fmt.Println("Output:", utils.SsoSlice([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

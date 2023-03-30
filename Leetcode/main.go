package main

import (
	"fmt"

	"leetcode.com/Leetcode/solutions"
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

func main() {
	PrintLine()
	SolveProblem()
	PrintLine()
	GetRandomProblem()
	PrintLine()
	// TestPattern()
}

func SolveProblem() {
	fmt.Println("Golang Leetcode...")
	solutions.Leetcode_Number_Of_Arrays()
}

func PrintLine() {
	fmt.Println("..............................................................................................................................................")
}

func GetRandomProblem() {
	Leetcode := new(types.Leetcode)
	Leetcode.SetTotal(2603)
	Leetcode.SetSolved(1, 2, 4, 5, 13, 14, 15, 20, 21, 23, 26, 27, 2571, 28, 35, 58, 67, 72, 101, 103, 104, 106, 109, 112, 121, 129, 142, 208, 211, 226, 290, 382, 427, 443, 502, 540, 567, 605)
	Leetcode.SetSolved(652, 783, 875, 9, 953, 958, 989, 1011, 989, 1071, 2582, 1137, 1162, 1207, 1319, 1345, 1466, 1470, 1472, 1523, 1539, 1675, 2187, 2316, 2348, 2360, 2444, 2477, 2492, 2566)
	Leetcode.SetSolved(3, 2567, 912, 2570, 2574, 2575, 9, 1129, 2576, 2578, 2579, 2580, 2583, 2586, 2587, 2588, 2591, 2592, 2595, 2596, 2597, 2600, 2601, 2602, 64, 983, 219, 1402, 2390, 768)
	Leetcode.SetSolved(769, 1944, 87, 2145, 601)

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problems has been solved in Leetcode.")
	fmt.Println("Today, Number of Leetcode Problem is:", Leetcode.GetRandom())
}

func TestPattern() {
	fmt.Println("Input: nums = [0,0,0,2,0,0]")
	fmt.Println("Output:", utils.Sslice([]int{0, 0, 0, 2, 0, 0}))
	fmt.Println("Input: grid = [[1,0,0],[0,0,0],[0,0,0]]")
	fmt.Println("Output:", utils.SsoSlice([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

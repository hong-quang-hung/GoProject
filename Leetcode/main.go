package main

import (
	"fmt"
	"regexp"

	"leetcode.com/Leetcode/solutions"
	"leetcode.com/Leetcode/types"
)

var debugProblem int
var totalProbem int

func init() {
	debugProblem = 2
	totalProbem = 2638
}

func main() {
	PrintLine()
	LeetcodeDebug()
	PrintLine()
	LeetcodeInfomation()
	PrintLine()
}

func PrintLine() {
	fmt.Println("..............................................................................................................................................")
}

func LeetcodeDebug() {
	fmt.Printf("Leetcode Debug Start...\n")
	solutions.Leetcode_Debug(debugProblem)
	fmt.Printf("Leetcode Debug End.\n")
}

func LeetcodeInfomation() {
	Leetcode := new(types.Leetcode)
	Leetcode.SetTotal(totalProbem)
	Leetcode.SetSolved(1, 2, 4, 5, 13, 14, 15, 1444, 21, 23, 26, 27, 2571, 28, 35, 58, 67, 72, 101, 103, 104, 106, 109, 112, 121, 129, 142, 208, 211, 226, 290, 382, 427, 443, 502, 540, 567, 605)
	Leetcode.SetSolved(652, 783, 875, 769, 953, 958, 989, 1011, 989, 1071, 2582, 1137, 1162, 1207, 1319, 1345, 1466, 1470, 1472, 1523, 1539, 1675, 2187, 2316, 2348, 2360, 2444, 2477, 2492, 2566)
	Leetcode.SetSolved(3, 2567, 912, 2570, 2574, 2575, 9, 1129, 2576, 2578, 2579, 2580, 2583, 2586, 2587, 2588, 2591, 2592, 2595, 2596, 2597, 2600, 2601, 2602, 64, 983, 219, 1402, 2390, 768, 87)
	Leetcode.SetSolved(769, 1944, 2145, 601, 20, 704, 2605, 2606, 2610, 32, 2609, 2300, 2306, 45, 2547, 881, 1954, 925, 2405, 1584, 2439, 1254, 1020, 1857, 133, 585, 1047, 299, 71)

	Leetcode.SetSubmit(1, 3, 4, 5, 2601, 27, 768, 769, 1944, 2145, 601, 20, 704, 2605, 2606, 2610, 32, 2609, 2300, 2306, 45, 2547, 881, 1954, 925, 2405, 1584, 2439, 1254, 2574, 103, 2575, 20, 13)
	Leetcode.SetSubmit(112, 64, 2439, 2578, 87, 704, 45, 2602, 35, 2586, 76, 1020, 226, 2596, 2492, 652, 1523, 9, 2610, 58, 1319, 783, 28, 133, 567, 2316, 2587, 2576, 2597, 129, 1047, 1470, 1466)
	Leetcode.SetSubmit(2586, 72, 605, 299, 989, 2187, 427, 1207, 2390, 2570, 1444, 2588, 2477, 67, 101, 1162, 2580, 2600, 211, 958, 71, 1345, 2583, 21, 106, 585, 2592, 1071, 443, 2)

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problem(s) has been solved in Leetcode.", "Golang is", Leetcode.Submit(), "submit(s).")
	// fmt.Println("Today, Number of Leetcode Problem is:", Leetcode.PickProblem())
	fmt.Println("Problem is", Leetcode.PickSolution(), "hasn't been submit.")
}

func TestRegexGolang() {
	re := regexp.MustCompile(`(\[[^\[]*\])`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[0],[],[1,112,-33,54, 45]", -1))

	re = regexp.MustCompile(`^\[(.+)\]$`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[1,112,-33,54, 45]", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[]", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[ ]", -1))
}

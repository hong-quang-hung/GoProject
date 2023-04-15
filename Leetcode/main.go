package main

import (
	"fmt"
	"regexp"

	"leetcode.com/Leetcode/solutions"
	"leetcode.com/Leetcode/types"
)

var _PROBLEM_DEBUG_ int
var _PROBLEM_TOTAL_ int

func init() {
	// _PROBLEM_DEBUG_ = 2218
	_PROBLEM_DEBUG_ = 6335
	_PROBLEM_TOTAL_ = 2642
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
	solutions.Leetcode_Debug(_PROBLEM_DEBUG_)
	fmt.Printf("Leetcode Debug End.\n")
}

func LeetcodeInfomation() {
	Leetcode := new(types.Leetcode)
	Leetcode.SetTotal(_PROBLEM_TOTAL_)

	Leetcode.SetSolved(1, 2, 4, 5, 13, 14, 15, 1444, 21, 23, 26, 27, 2571, 28, 35, 58, 67, 72, 104, 106, 109, 112, 121, 129, 142, 208, 211, 226, 290, 382, 427, 443, 502, 540, 567, 605)
	Leetcode.SetSolved(783, 875, 769, 953, 989, 1011, 989, 1071, 2582, 1137, 1162, 1207, 1319, 1345, 1466, 1470, 1472, 1523, 1539, 1675, 2187, 2316, 2348, 2360, 2444, 2477, 2492, 2566)
	Leetcode.SetSolved(3, 2567, 570, 2574, 2575, 9, 1129, 2576, 2578, 2579, 2580, 2583, 2586, 2587, 2588, 2591, 2592, 2595, 2596, 2597, 2600, 2601, 2602, 983, 219, 1402, 2390, 768, 87)
	Leetcode.SetSolved(769, 1944, 2145, 601, 20, 64, 2605, 2606, 2610, 32, 2609, 2300, 2306, 45, 2547, 881, 1954, 925, 2405, 1584, 2439, 1254, 1020, 1857, 133, 585, 1047, 299, 71, 946)
	Leetcode.SetSolved(652, 912, 101, 516, 958, 103, 704, 2399, 1572, 1925, 2639, 2640)

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problem(s) has been solved in Leetcode.")
	// fmt.Println("Today, Number of Leetcode Problem is:", Leetcode.PickProblem())
}

func TestRegexGolang() {
	re := regexp.MustCompile(`(\[[^\[]*\])`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[0],[],[1,112,-33,54, 45]", -1))

	re = regexp.MustCompile(`^\[(.+)\]$`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[1,112,-33,54, 45]", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[]", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("[ ]", -1))
}

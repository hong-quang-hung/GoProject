package main

import (
	"fmt"
	"regexp"

	"leetcode.com/Leetcode/solutions"
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

var _PROBLEM_DEBUG_ int
var _PROBLEM_TOTAL_ int

func init() {
	_PROBLEM_DEBUG_ = 102275 - solutions.Normalize

	_PROBLEM_TOTAL_ = 2689
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

	Leetcode.SetSolved(0, 10, 20, 40, 50, 290, 540, 1470, 2130, 2140, 2390, 2610, 2640, 2660, 2680)
	Leetcode.SetSolved(1, 11, 21, 71, 101, 121, 141, 211, 231, 401, 1011, 1071, 1721, 2571, 226, 382, 427, 443, 567, 605, 2678)
	Leetcode.SetSolved(2, 12, 22, 32, 72, 112, 142, 202, 342, 502, 652, 912, 1312, 1572, 2582, 1162, 1345, 1466, 1472, 1523, 1539, 1675, 2187, 2316, 2348, 2360, 2444, 2477, 2492, 2652, 2566)
	Leetcode.SetSolved(3, 13, 23, 103, 263, 783, 953, 983, 2574, 2575, 2576, 2578, 2579, 2580, 2583, 2586, 2587, 2588, 2591, 2592, 2595, 2596, 2597, 2600, 2601, 2602, 1402, 768, 2679)
	Leetcode.SetSolved(4, 14, 24, 54, 64, 104, 904, 1444, 1944, 2044, 2145, 601, 2605, 2606, 2609, 2300, 2306, 881, 1954, 925, 2405, 1584, 2439, 1254, 1020, 1857, 133, 585, 1047, 299, 946, 2682)
	Leetcode.SetSolved(5, 15, 35, 45, 875, 1805, 516, 958, 704, 2399, 1925, 2639, 2641, 2218, 2643, 2644, 2645, 198, 1639, 1431, 1348, 1768, 509, 1372, 570, 2570, 1049, 662, 879, 213, 2685, 2683)
	Leetcode.SetSolved(6, 16, 26, 36, 106, 216, 326, 2646, 2520, 2466, 2656, 849, 2653, 2651, 204, 2654, 118, 1046, 309, 2336, 258, 1582, 319, 264, 279, 313, 1201, 2684)
	Leetcode.SetSolved(7, 17, 27, 67, 77, 87, 337, 377, 1137, 1207, 1557, 2547, 2567, 1697, 2657, 2661, 2659, 1579, 1491, 1416, 1822, 2215, 1964, 1456, 2662, 1498, 1965, 2658, 2673, 2670, 2671, 2672, 1035)
	Leetcode.SetSolved(8, 18, 28, 38, 58, 208, 438, 2038, 78, 2275)
	Leetcode.SetSolved(9, 19, 39, 59, 109, 119, 129, 219, 649, 769, 839, 929, 989, 1129, 1319, 1799, 1979, 2419)

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problem(s) has been solved in Leetcode.")
	fmt.Println("Today, Number of Leetcode Problem is:", Leetcode.PickProblem())

	ShowHasNotSubmited(Leetcode)
}

func TestRegexGolang() {
	re := regexp.MustCompile(`[-]`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("42", -1))
}

func TestTreeNode() {
	fmt.Println("root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]")
	root1 := utils.S2TreeNode("[1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]")
	root2 := types.LazyNodeAll(3, types.LazyNodeAll(1, nil, types.LazyNode(4)), types.LazyNode(2))
	fmt.Println(utils.STreeNode(root1))
	fmt.Println(utils.STreeNode(root2))
}

func ShowHasNotSubmited(L *types.Leetcode) {
	for i := 0; i < _PROBLEM_TOTAL_; i++ {
		if L.IsSolved(i) && !solutions.Leetcode_Check_Golang_Solution(i+1) {
			fmt.Println(i+1, "hasn't submit solution with Golang language.")
		}
	}
}

package main

import (
	"fmt"

	"leetcode.com/Leetcode/solutions"
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

var (
	_PROBLEM_DEBUG_ int
	_PROBLEM_TOTAL_ int
	_PROBLEM_GROUP_ int
)

func init() {
	_PROBLEM_DEBUG_ = 101998 - solutions.Normalize
	_PROBLEM_DEBUG_ = 101514 - solutions.Normalize

	_PROBLEM_GROUP_ = 10
	_PROBLEM_TOTAL_ = 2759
}

func main() {
	PrintLine()
	LeetcodeDebug()
	PrintLine()
	LeetcodeInformation()
	PrintLine()
}

func PrintLine() {
	fmt.Println("..............................................................................................................................................")
}

func LeetcodeDebug() {
	fmt.Printf("Leetcode Debug Start...\n")
	solutions.Leetcode_Debug(_PROBLEM_DEBUG_)
	fmt.Printf("Leetcode Debug End. The question %d belongs to group %dth.\n", _PROBLEM_DEBUG_, (_PROBLEM_DEBUG_)%_PROBLEM_GROUP_)
}

func LeetcodeInformation() {
	Leetcode := new(types.Leetcode)
	Leetcode.SetTotal(_PROBLEM_TOTAL_)

	Leetcode.SetSolved(0, 10, 20, 40, 50, 70, 80, 290, 530, 540, 570, 1020, 1140, 1470, 2090, 2130, 2140, 2300, 2360, 2390, 2520, 2570, 2580, 2600, 2610, 2620, 2640, 2660, 2670, 2680, 2700, 2710, 2730, 2740)
	Leetcode.SetSolved(1, 11, 21, 71, 101, 121, 141, 151, 211, 231, 401, 601, 881, 1011, 1071, 1091, 1161, 1201, 1351, 1431, 1491, 1721, 2101, 2571, 2591, 2601, 2621, 2631, 2641, 2651, 2661, 2671, 2711, 2721, 2731, 2741)
	Leetcode.SetSolved(2, 12, 22, 32, 62, 72, 112, 142, 202, 342, 382, 392, 502, 652, 662, 912, 1092, 1162, 1202, 1232, 1312, 1372, 1402, 1472, 1502, 1572, 1582, 1732, 1802, 1822, 2352, 2462, 2492, 2542, 2582, 2592, 2602, 2622, 2632, 2652, 2662, 2672, 2682, 2712, 2722, 2732, 2742)
	Leetcode.SetSolved(3, 13, 23, 103, 133, 213, 263, 283, 313, 373, 443, 583, 643, 703, 783, 933, 953, 983, 1143, 1493, 1523, 1603, 2583, 2623, 2633, 2643, 2653, 2673, 2683, 2693, 2703, 2723, 2733)
	Leetcode.SetSolved(4, 14, 24, 54, 64, 104, 204, 264, 334, 394, 704, 714, 724, 744, 904, 934, 1004, 1254, 1444, 1514, 1584, 1944, 1954, 1964, 2044, 2444, 2574, 2634, 2644, 2654, 2684, 2694, 2704, 2724, 2734)
	Leetcode.SetSolved(5, 15, 25, 35, 45, 345, 585, 605, 705, 735, 785, 875, 925, 1035, 1345, 1675, 1805, 1925, 1965, 2145, 2215, 2275, 2405, 2575, 2595, 2605, 2625, 2635, 2645, 2665, 2675, 2685, 2695, 2705, 2715, 2725, 2735)
	Leetcode.SetSolved(6, 16, 26, 36, 46, 106, 216, 226, 326, 516, 946, 956, 1046, 1146, 1376, 1396, 1406, 1416, 1456, 1466, 2306, 2316, 2336, 2646, 2466, 2566, 2576, 2586, 2596, 2606, 2626, 2636, 2656, 2666, 2676, 2696, 2706, 2716, 2726)
	Leetcode.SetSolved(7, 17, 27, 67, 77, 87, 337, 347, 377, 427, 547, 567, 837, 1027, 1047, 1137, 1187, 1207, 1547, 1557, 1657, 1697, 1857, 2187, 2477, 2547, 2567, 2587, 2597, 2627, 2637, 2657, 2667, 2677, 2697, 2707, 2717, 2727)
	Leetcode.SetSolved(8, 18, 28, 38, 58, 78, 88, 118, 198, 208, 228, 238, 258, 438, 768, 958, 1318, 1348, 1498, 1768, 2038, 2218, 2328, 2348, 2448, 2578, 2588, 2618, 2628, 2648, 2658, 2678, 2698, 2708, 2718)
	Leetcode.SetSolved(9, 19, 39, 59, 109, 119, 129, 219, 279, 299, 309, 319, 389, 399, 509, 649, 769, 839, 849, 879, 929, 989, 1049, 1129, 1319, 1539, 1569, 1579, 1639, 1679, 1799, 1979, 2399, 2419, 2439, 2579, 2609, 2619, 2629, 2639, 2649, 2659, 2679, 2719, 2729, 2739)

	// defer ShowHasNotSubmited(Leetcode)
	// defer LeetcodeGroupBy(Leetcode, 20)

	fmt.Println("There are", Leetcode.Solved(), "/", Leetcode.Total(), "problem(s) has been solved in Leetcode.")
	// fmt.Println("Today, Number of Leetcode Problem is:", Leetcode.PickProblem())
}

func ShowHasNotSubmited(L *types.Leetcode) {
	for i := 0; i < _PROBLEM_TOTAL_; i++ {
		if L.IsSolved(i) && !solutions.Leetcode_Check_Golang_Solution(i+1) {
			fmt.Println(i+1, "hasn't submit solution with Golang language.")
		}
	}
}

func LeetcodeGroupBy(L *types.Leetcode, n int) {
	group := make([][]int, n)
	for i := 0; i < n; i++ {
		group[i] = []int{}
	}

	group[0] = append(group[0], 0)
	for i := -1; i < _PROBLEM_TOTAL_; i++ {
		if L.IsSolved(i) {
			group[(i+1)%n] = append(group[(i+1)%n], i+1)
		}
	}

	PrintLine()
	fmt.Println("Leetcode group by:", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%2d : %s\n", i, utils.Sslice(group[i]))
	}
}

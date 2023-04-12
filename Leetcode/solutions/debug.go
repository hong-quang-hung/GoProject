package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/solutions/easy"
	"leetcode.com/Leetcode/solutions/hard"
	"leetcode.com/Leetcode/solutions/medium"
	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

var debugFunc map[int]func()

func init() {
	debugFunc = make(map[int]func())
	// Easy
	debugFunc[1] = easy.Leetcode_Two_Sum
	debugFunc[9] = easy.LeetCode_Is_Palindrome
	debugFunc[20] = easy.Leetcode_Is_Valid
	debugFunc[21] = easy.Leetcode_Merge_Two_Lists
	debugFunc[26] = easy.Leetcode_Remove_Duplicates
	debugFunc[27] = easy.Leetcode_Remove_Element
	debugFunc[28] = easy.Leetcode_StrStr
	debugFunc[35] = easy.Leetcode_Search_Insert
	debugFunc[58] = easy.Leetcode_Length_Of_Last_Word
	debugFunc[67] = easy.Leetcode_Add_Binary
	debugFunc[101] = easy.Leetcode_Is_Symmetric
	debugFunc[104] = easy.LeetCode_Max_Depth
	debugFunc[121] = easy.LeetCode_Max_Profit
	debugFunc[219] = easy.Leetcode_Contains_Near_By_Duplicate
	debugFunc[290] = easy.Leetcode_Word_Pattern
	debugFunc[704] = easy.Leetcode_Search
	debugFunc[925] = easy.Leetcode_Is_Long_PressedName
	debugFunc[953] = easy.Leetcode_Is_Alien_Sorted
	debugFunc[1071] = easy.Leetcode_Gcd_Of_Strings
	debugFunc[1207] = easy.Leetcode_Unique_Occurrences
	debugFunc[1470] = easy.Leetcode_Shuffle
	debugFunc[2586] = easy.Leetcode_Vowel_Strings
	debugFunc[2595] = easy.Leetcode_Even_Odd_Bit
	debugFunc[2600] = easy.Leetcode_K_Items_With_Maximum_Sum
	debugFunc[2605] = easy.Leetcode_Min_Number
	debugFunc[2609] = easy.Leetcode_Find_The_Longest_Balanced_Substring
	// Medium
	debugFunc[3] = medium.Leetcode_Length_Of_Longest_Substring
	debugFunc[5] = medium.Leetcode_Longest_Palindrome
	debugFunc[45] = medium.Leetcode_Jumb_Game_II
	debugFunc[64] = medium.Leetcode_Min_Path_Sum
	debugFunc[71] = medium.Leetcode_Simplify_Path
	debugFunc[103] = medium.Leetcode_Zigzag_Level_Order
	debugFunc[106] = medium.Leetcode_Build_Trees
	debugFunc[129] = medium.Leetcode_Sum_Numbers
	debugFunc[142] = medium.Leetcode_Detect_Cycle
	debugFunc[211] = medium.Leetcode_Word_Dictionary
	debugFunc[769] = medium.Leetcode_Max_Chunks_To_Sorted
	debugFunc[881] = medium.Leetcode_Num_Rescue_Boats
	debugFunc[875] = medium.Leetcode_Min_Eating_Speed
	debugFunc[912] = medium.Leetcode_Sort_Array
	debugFunc[958] = medium.Leetcode_Is_Complete_Tree
	debugFunc[983] = medium.Leetcode_Min_Cost_Tickets
	debugFunc[1129] = medium.Leetcode_Shortest_Alternating_Paths
	debugFunc[1162] = medium.Leetcode_Max_Distance
	debugFunc[1254] = medium.Leetcode_Closed_Island
	debugFunc[1319] = medium.Leetcode_Make_Connected
	debugFunc[1584] = medium.Leetcode_minCost_Connect_Points
	debugFunc[1954] = medium.Leetcode_Minimum_Perimeter
	debugFunc[2145] = medium.Leetcode_Number_Of_Arrays
	debugFunc[2300] = medium.Leetcode_Successful_Pairs
	debugFunc[2316] = medium.Leetcode_Count_Pairs
	debugFunc[2348] = medium.Leetcode_Zero_Filled_Subarray
	debugFunc[2390] = medium.Leetcode_Remove_Stars
	debugFunc[2405] = medium.Leetcode_Partition_String
	debugFunc[2439] = medium.Leetcode_Minimize_Array_Value
	debugFunc[2444] = medium.Leetcode_Count_Subarrays
	debugFunc[2492] = medium.Leetcode_Min_Score
	debugFunc[2579] = medium.Leetcode_Colored_Cells
	debugFunc[2583] = medium.Leetcode_Kth_Largest_Level_Sum
	debugFunc[2587] = medium.Leetcode_Max_Score
	debugFunc[2588] = medium.Leetcode_Beautiful_Subarrays
	debugFunc[2592] = medium.Leetcode_Maximize_Greatness
	debugFunc[2597] = medium.Leetcode_Beautiful_Subsets
	debugFunc[2601] = medium.Leetcode_Prime_Sub_Operation
	debugFunc[2602] = medium.Leetcode_Min_Operations
	debugFunc[2606] = medium.Leetcode_Maximum_Cost_Substring
	// Hard
	debugFunc[4] = hard.Leetcode_Find_Median_Sorted_Arrays
	debugFunc[32] = hard.Leetcode_Longest_Valid_Parentheses
	debugFunc[87] = hard.Leetcode_Is_Scramble
	debugFunc[768] = hard.Leetcode_Max_Chunks_To_Sorted_II
	debugFunc[1402] = hard.Leetcode_Max_Satisfaction
	debugFunc[1444] = hard.Leetcode_Ways
	debugFunc[1675] = hard.Leetcode_Minimum_Deviation
	debugFunc[1857] = hard.Leetcode_Largest_Path_Value
	debugFunc[1944] = hard.Leetcode_Can_See_Persons_Count
	debugFunc[2306] = hard.Leetcode_Distinct_Names
	debugFunc[2360] = hard.Leetcode_Longest_Cycle
	debugFunc[2547] = hard.Leetcode_Min_Cost
	// Other
	debugFunc[585] = Leetcode_SQL
	debugFunc[601] = Leetcode_SQL
}

func Leetcode_Debug(problem int) {
	if invoke, c := debugFunc[problem]; c {
		invoke()
	} else {
		Leetcode_Empty(problem)
	}
}

func Leetcode_Empty(problem int) {
	fmt.Printf("The problem %d hasn't been solved yet!\n", problem)
}

func Leetcode_SQL() {
	fmt.Printf("This is SQL solution!\n")
}

func Leetcode_Find_Matrix() {
	fmt.Println("Input: nums = [1,3,4,1,2,3,1]")
	fmt.Println("Output:", findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output:", findMatrix([]int{1, 2, 3, 4}))
}

func Leetcode_Split_Num() {
	fmt.Println("Input: num = 4325")
	fmt.Println("Output:", splitNum(4325))
}

func Leetcode_Left_Rigth_Difference() {
	fmt.Println("Input: nums = [10,4,8,3]")
	fmt.Println("Output:", leftRigthDifference(utils.S2SliceInt("[10,4,8,3]")))
}

func Leetcode_Divisibility_Array() {
	fmt.Println("Input: word = '998244353', m = 3")
	fmt.Println("Output:", divisibilityArray("998244353", 3))
}

func Leetcode_Num_Enclaves() {
	fmt.Println("Input: grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]")
	fmt.Println("Output:", numEnclaves(utils.S2SoSliceInt("[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]")))
	fmt.Println("Input: grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]")
	fmt.Println("Output:", numEnclaves(utils.S2SoSliceInt("[[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]")))
}

func Leetcode_Invert_Tree() {
	fmt.Println("Input: root = [4,2,7,1,3,6,9]")
	fmt.Println("Output:", invertTree(types.LazyNodeAll(4, types.LazyNodeValue(2, 1, 3), types.LazyNodeValue(7, 6, 9))))
}

func Leetcode_Check_Valid_Grid() {
	fmt.Println("Input: grid = [[8,3,6],[5,0,1],[2,7,4]]")
	fmt.Println("Output:", checkValidGrid(utils.S2SoSliceInt("[[8,3,6],[5,0,1],[2,7,4]]")))
}

func Leetcode_Find_Duplicate_Subtrees() {
	fmt.Println("Input: root = [1,2,3,4,null,2,4,null,null,4]")
	fmt.Println("Output:", findDuplicateSubtrees(types.LazyNodeAll(1, types.LazyNodeAll(2, types.LazyNode(4), nil), types.LazyNodeAll(3, types.LazyNodeAll(2, types.LazyNode(4), nil), types.LazyNode(4)))))
}

func Leetcode_Min_Diff_In_BST() {
	fmt.Println("Input: root = [4,2,6,1,3]")
	fmt.Println("Output:", minDiffInBST(types.LazyNodeAll(4, types.LazyNodeValue(2, 1, 3), types.LazyNode(6))))
}

func Leetcode_Clone_Graph() {
	fmt.Println("Input: adjList = [[2,4],[1,3],[2,4],[1,3]]")
	adjList1 := new(types.Node)
	adjList2 := new(types.Node)
	adjList3 := new(types.Node)
	adjList4 := new(types.Node)

	adjList1.Val = 1
	adjList1.Neighbors = []*types.Node{adjList2, adjList4}
	adjList2.Val = 2
	adjList2.Neighbors = []*types.Node{adjList1, adjList3}
	adjList3.Val = 3
	adjList3.Neighbors = []*types.Node{adjList2, adjList4}
	adjList4.Val = 4
	adjList3.Neighbors = []*types.Node{adjList1, adjList3}
	fmt.Println("Output:", cloneGraph(adjList1))
}

func Leetcode_Check_Inclusion() {
	fmt.Println("Input: s1 = 'ab', s2 = 'eidbaooo'")
	fmt.Println("Output:", checkInclusion("ab", "eidbaooo"))
}

func Leetcode_Max_Num_Of_Marked_Indices() {
	fmt.Println("Input: nums = [3,5,2,4]")
	fmt.Println("Output:", maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
}

func Leetcode_Remove_Adjacent_Duplicates() {
	fmt.Println("Input: s = 'azxxzy'")
	fmt.Println("Output:", removeAdjacentDuplicates("azxxzy"))
}

func Leetcode_Min_Reorder() {
	fmt.Println("Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]")
	fmt.Println("Output:", minReorder(6, utils.S2SoSliceInt("[[0,1],[1,3],[2,3],[4,0],[4,5]]")))
}

func Leetcode_Min_Distance() {
	fmt.Println("Input: word1 = 'intention', word2 = 'execution'")
	fmt.Println("Output:", minDistance("intention", "execution"))
}

func Leetcode_Can_Place_Flowers() {
	fmt.Println("Input: flowerbed = [1,0,0,0,1], n = 1")
	fmt.Println("Output:", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}

func Leetcode_Get_Hint() {
	fmt.Println("Input: secret = '1234', guess = '0111'")
	fmt.Println("Output:", getHint("1234", "0111"))
	fmt.Println("Input: secret = '1122', guess = '2211'")
	fmt.Println("Output:", getHint("1122", "2211"))
}

func Leetcode_Add_To_Array_Form() {
	fmt.Println("Input: num = [1,2,0,0], k = 34")
	fmt.Println("Output:", addToArrayForm([]int{1, 2, 0, 0}, 34))
}

func Leetcode_Minimum_Time() {
	fmt.Println("Input: time = [1,2,3], totalTrips = 5")
	fmt.Println("Output:", minimumTime([]int{1, 2, 3}, 5))
}

func Leetcode_Construct() {
	fmt.Println("Input: grid = [[0,1],[1,0]]")
	fmt.Println("Output:", construct(utils.S2SoSliceInt("[[0,1],[1,0]]")))
}

func Leetcode_Merge_Arrays() {
	fmt.Println("Input: nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]")
	fmt.Println("Output:", mergeArrays(utils.S2SoSliceInt("[[1,2],[2,3],[4,5]]"), utils.S2SoSliceInt("[[1,4],[3,2],[4,1]]")))
}

func Leetcode_Minimum_Fuel_Cost() {
	fmt.Println("Input: roads = [[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]], seats = 2")
	fmt.Println("Output:", minimumFuelCost(utils.S2SoSliceInt("[[3,1],[3,2],[1,0],[0,4],[0,5],[4,6]]"), 2))
}

func Leetcode_Count_Ways() {
	fmt.Println("Input: ranges = [[6,10],[5,15]]")
	fmt.Println("Output:", countWays(utils.S2SoSliceInt("[[6,10],[5,15]]")))
}

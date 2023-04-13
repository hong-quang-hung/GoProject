package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/solutions/easy"
	"leetcode.com/Leetcode/solutions/hard"
	"leetcode.com/Leetcode/solutions/medium"
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
	debugFunc[226] = easy.Leetcode_Invert_Tree
	debugFunc[290] = easy.Leetcode_Word_Pattern
	debugFunc[605] = easy.Leetcode_Can_Place_Flowers
	debugFunc[704] = easy.Leetcode_Search
	debugFunc[783] = easy.Leetcode_Min_Diff_In_BST
	debugFunc[925] = easy.Leetcode_Is_Long_PressedName
	debugFunc[953] = easy.Leetcode_Is_Alien_Sorted
	debugFunc[989] = easy.Leetcode_Add_To_Array_Form
	debugFunc[1047] = easy.Leetcode_Remove_Adjacent_Duplicates
	debugFunc[1071] = easy.Leetcode_Gcd_Of_Strings
	debugFunc[1207] = easy.Leetcode_Unique_Occurrences
	debugFunc[1470] = easy.Leetcode_Shuffle
	debugFunc[1539] = easy.Leetcode_Find_Kth_Positive
	debugFunc[2566] = easy.Leetcode_Min_Max_Difference
	debugFunc[2570] = easy.Leetcode_Merge_Arrays
	debugFunc[2574] = easy.Leetcode_Left_Rigth_Difference
	debugFunc[2578] = easy.Leetcode_Split_Num
	debugFunc[2586] = easy.Leetcode_Vowel_Strings
	debugFunc[2591] = easy.Leetcode_Dist_Money
	debugFunc[2595] = easy.Leetcode_Even_Odd_Bit
	debugFunc[2600] = easy.Leetcode_K_Items_With_Maximum_Sum
	debugFunc[2605] = easy.Leetcode_Min_Number
	debugFunc[2609] = easy.Leetcode_Find_The_Longest_Balanced_Substring
	// Medium
	debugFunc[2] = medium.Leetcode_Add_Two_Numbers
	debugFunc[3] = medium.Leetcode_Length_Of_Longest_Substring
	debugFunc[5] = medium.Leetcode_Longest_Palindrome
	debugFunc[45] = medium.Leetcode_Jumb_Game_II
	debugFunc[64] = medium.Leetcode_Min_Path_Sum
	debugFunc[71] = medium.Leetcode_Simplify_Path
	debugFunc[103] = medium.Leetcode_Zigzag_Level_Order
	debugFunc[106] = medium.Leetcode_Build_Trees
	debugFunc[129] = medium.Leetcode_Sum_Numbers
	debugFunc[133] = medium.Leetcode_Clone_Graph
	debugFunc[142] = medium.Leetcode_Detect_Cycle
	debugFunc[211] = medium.Leetcode_Word_Dictionary
	debugFunc[299] = medium.Leetcode_Get_Hint
	debugFunc[382] = medium.Leetcode_Constructor
	debugFunc[427] = medium.Leetcode_Construct
	debugFunc[443] = medium.Leetcode_Compress
	debugFunc[567] = medium.Leetcode_Check_Inclusion
	debugFunc[652] = medium.Leetcode_Find_Duplicate_Subtrees
	debugFunc[769] = medium.Leetcode_Max_Chunks_To_Sorted
	debugFunc[881] = medium.Leetcode_Num_Rescue_Boats
	debugFunc[875] = medium.Leetcode_Min_Eating_Speed
	debugFunc[912] = medium.Leetcode_Sort_Array
	debugFunc[958] = medium.Leetcode_Is_Complete_Tree
	debugFunc[983] = medium.Leetcode_Min_Cost_Tickets
	debugFunc[1020] = medium.Leetcode_Num_Enclaves
	debugFunc[1129] = medium.Leetcode_Shortest_Alternating_Paths
	debugFunc[1162] = medium.Leetcode_Max_Distance
	debugFunc[1254] = medium.Leetcode_Closed_Island
	debugFunc[1319] = medium.Leetcode_Make_Connected
	debugFunc[1466] = medium.Leetcode_Min_Reorder
	debugFunc[1584] = medium.Leetcode_minCost_Connect_Points
	debugFunc[1954] = medium.Leetcode_Minimum_Perimeter
	debugFunc[2145] = medium.Leetcode_Number_Of_Arrays
	debugFunc[2187] = medium.Leetcode_Minimum_Time
	debugFunc[2300] = medium.Leetcode_Successful_Pairs
	debugFunc[2316] = medium.Leetcode_Count_Pairs
	debugFunc[2348] = medium.Leetcode_Zero_Filled_Subarray
	debugFunc[2390] = medium.Leetcode_Remove_Stars
	debugFunc[2405] = medium.Leetcode_Partition_String
	debugFunc[2439] = medium.Leetcode_Minimize_Array_Value
	debugFunc[2444] = medium.Leetcode_Count_Subarrays
	debugFunc[2477] = medium.Leetcode_Minimum_Fuel_Cost
	debugFunc[2492] = medium.Leetcode_Min_Score
	debugFunc[2567] = medium.Leetcode_Minimize_Sum
	debugFunc[2571] = medium.Leetcode_Min_Operations
	debugFunc[2575] = medium.Leetcode_Divisibility_Array
	debugFunc[2576] = medium.Leetcode_Max_Num_Of_Marked_Indices
	debugFunc[2579] = medium.Leetcode_Colored_Cells
	debugFunc[2580] = medium.Leetcode_Count_Ways
	debugFunc[2583] = medium.Leetcode_Kth_Largest_Level_Sum
	debugFunc[2587] = medium.Leetcode_Max_Score
	debugFunc[2588] = medium.Leetcode_Beautiful_Subarrays
	debugFunc[2592] = medium.Leetcode_Maximize_Greatness
	debugFunc[2596] = medium.Leetcode_Check_Valid_Grid
	debugFunc[2597] = medium.Leetcode_Beautiful_Subsets
	debugFunc[2601] = medium.Leetcode_Prime_Sub_Operation
	debugFunc[2602] = medium.Leetcode_Min_Operations_I
	debugFunc[2606] = medium.Leetcode_Maximum_Cost_Substring
	debugFunc[2610] = medium.Leetcode_Find_Matrix
	// Hard
	debugFunc[4] = hard.Leetcode_Find_Median_Sorted_Arrays
	debugFunc[23] = hard.Leetcode_Merge_K_Lists
	debugFunc[32] = hard.Leetcode_Longest_Valid_Parentheses
	debugFunc[72] = hard.Leetcode_Min_Distance
	debugFunc[87] = hard.Leetcode_Is_Scramble
	debugFunc[768] = hard.Leetcode_Max_Chunks_To_Sorted_II
	debugFunc[1345] = hard.Leetcode_Min_Jumps
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

package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/solutions/easy"
	"leetcode.com/Leetcode/solutions/hard"
	"leetcode.com/Leetcode/solutions/medium"
)

type solutions map[int]func()

// Format problem th to int
const normalize int = 100000

// problem th convert to problem + normalize
var debugFunc solutions

func init() {
	debugFunc = make(map[int]func())
	// Easy
	debugFunc[100001] = easy.Leetcode_Two_Sum
	debugFunc[100009] = easy.LeetCode_Is_Palindrome
	debugFunc[100014] = easy.Leetcode_Longest_Common_Prefix
	debugFunc[100020] = easy.Leetcode_Is_Valid
	debugFunc[100021] = easy.Leetcode_Merge_Two_Lists
	debugFunc[100026] = easy.Leetcode_Remove_Duplicates
	debugFunc[100027] = easy.Leetcode_Remove_Element
	debugFunc[100028] = easy.Leetcode_StrStr
	debugFunc[100035] = easy.Leetcode_Search_Insert
	debugFunc[100058] = easy.Leetcode_Length_Of_Last_Word
	debugFunc[100067] = easy.Leetcode_Add_Binary
	debugFunc[100101] = easy.Leetcode_Is_Symmetric
	debugFunc[100104] = easy.LeetCode_Max_Depth
	debugFunc[100121] = easy.LeetCode_Max_Profit
	debugFunc[100219] = easy.Leetcode_Contains_Near_By_Duplicate
	debugFunc[100226] = easy.Leetcode_Invert_Tree
	debugFunc[100290] = easy.Leetcode_Word_Pattern
	debugFunc[100605] = easy.Leetcode_Can_Place_Flowers
	debugFunc[100704] = easy.Leetcode_Search
	debugFunc[100783] = easy.Leetcode_Min_Diff_In_BST
	debugFunc[100925] = easy.Leetcode_Is_Long_PressedName
	debugFunc[100953] = easy.Leetcode_Is_Alien_Sorted
	debugFunc[100989] = easy.Leetcode_Add_To_Array_Form
	debugFunc[101047] = easy.Leetcode_Remove_Adjacent_Duplicates
	debugFunc[101071] = easy.Leetcode_Gcd_Of_Strings
	debugFunc[101137] = easy.Leetcode_Tribonacci
	debugFunc[101207] = easy.Leetcode_Unique_Occurrences
	debugFunc[101470] = easy.Leetcode_Shuffle
	debugFunc[101539] = easy.Leetcode_Find_Kth_Positive
	debugFunc[101572] = easy.Leetcode_Diagonal_Sum
	debugFunc[102399] = easy.Leetcode_Check_Distances
	debugFunc[102566] = easy.Leetcode_Min_Max_Difference
	debugFunc[102570] = easy.Leetcode_Merge_Arrays
	debugFunc[102574] = easy.Leetcode_Left_Rigth_Difference
	debugFunc[102578] = easy.Leetcode_Split_Num
	debugFunc[102582] = easy.Leetcode_Pass_The_Pillow
	debugFunc[102586] = easy.Leetcode_Vowel_Strings
	debugFunc[102591] = easy.Leetcode_Dist_Money
	debugFunc[102595] = easy.Leetcode_Even_Odd_Bit
	debugFunc[102600] = easy.Leetcode_K_Items_With_Maximum_Sum
	debugFunc[102605] = easy.Leetcode_Min_Number
	debugFunc[102609] = easy.Leetcode_Find_The_Longest_Balanced_Substring
	// Medium
	debugFunc[100002] = medium.Leetcode_Add_Two_Numbers
	debugFunc[100003] = medium.Leetcode_Length_Of_Longest_Substring
	debugFunc[100005] = medium.Leetcode_Longest_Palindrome
	debugFunc[100015] = medium.Leetcode_Three_Sum
	debugFunc[100045] = medium.Leetcode_Jumb_Game_II
	debugFunc[100064] = medium.Leetcode_Min_Path_Sum
	debugFunc[100071] = medium.Leetcode_Simplify_Path
	debugFunc[100103] = medium.Leetcode_Zigzag_Level_Order
	debugFunc[100106] = medium.Leetcode_Build_Trees
	debugFunc[100109] = medium.Leetcode_Sorted_List_To_BST
	debugFunc[100129] = medium.Leetcode_Sum_Numbers
	debugFunc[100133] = medium.Leetcode_Clone_Graph
	debugFunc[100142] = medium.Leetcode_Detect_Cycle
	debugFunc[100208] = medium.Leetcode_Trie_Constructor
	debugFunc[100211] = medium.Leetcode_Word_Dictionary
	debugFunc[100299] = medium.Leetcode_Get_Hint
	debugFunc[100382] = medium.Leetcode_Constructor
	debugFunc[100427] = medium.Leetcode_Construct
	debugFunc[100443] = medium.Leetcode_Compress
	debugFunc[100516] = medium.Leetcode_Longest_Palindrome_Subseq
	debugFunc[100540] = medium.Leetcode_Single_Non_Duplicate
	debugFunc[100567] = medium.Leetcode_Check_Inclusion
	debugFunc[100652] = medium.Leetcode_Find_Duplicate_Subtrees
	debugFunc[100769] = medium.Leetcode_Max_Chunks_To_Sorted
	debugFunc[100881] = medium.Leetcode_Num_Rescue_Boats
	debugFunc[100875] = medium.Leetcode_Min_Eating_Speed
	debugFunc[100912] = medium.Leetcode_Sort_Array
	debugFunc[100946] = medium.Leetcode_Validate_Stack_Sequences
	debugFunc[100958] = medium.Leetcode_Is_Complete_Tree
	debugFunc[100983] = medium.Leetcode_Min_Cost_Tickets
	debugFunc[101011] = medium.Leetcode_Ship_Within_Days
	debugFunc[101020] = medium.Leetcode_Num_Enclaves
	debugFunc[101129] = medium.Leetcode_Shortest_Alternating_Paths
	debugFunc[101162] = medium.Leetcode_Max_Distance
	debugFunc[101254] = medium.Leetcode_Closed_Island
	debugFunc[101319] = medium.Leetcode_Make_Connected
	debugFunc[101466] = medium.Leetcode_Min_Reorder
	debugFunc[101472] = medium.Leetcode_Design_Browser_History
	debugFunc[101584] = medium.Leetcode_minCost_Connect_Points
	debugFunc[101954] = medium.Leetcode_Minimum_Perimeter
	debugFunc[102145] = medium.Leetcode_Number_Of_Arrays
	debugFunc[102187] = medium.Leetcode_Minimum_Time
	debugFunc[102300] = medium.Leetcode_Successful_Pairs
	debugFunc[102316] = medium.Leetcode_Count_Pairs
	debugFunc[102348] = medium.Leetcode_Zero_Filled_Subarray
	debugFunc[102390] = medium.Leetcode_Remove_Stars
	debugFunc[102405] = medium.Leetcode_Partition_String
	debugFunc[102439] = medium.Leetcode_Minimize_Array_Value
	debugFunc[102477] = medium.Leetcode_Minimum_Fuel_Cost
	debugFunc[102492] = medium.Leetcode_Min_Score
	debugFunc[102567] = medium.Leetcode_Minimize_Sum
	debugFunc[102571] = medium.Leetcode_Min_Operations
	debugFunc[102575] = medium.Leetcode_Divisibility_Array
	debugFunc[102576] = medium.Leetcode_Max_Num_Of_Marked_Indices
	debugFunc[102579] = medium.Leetcode_Colored_Cells
	debugFunc[102580] = medium.Leetcode_Count_Ways
	debugFunc[102583] = medium.Leetcode_Kth_Largest_Level_Sum
	debugFunc[102587] = medium.Leetcode_Max_Score
	debugFunc[102588] = medium.Leetcode_Beautiful_Subarrays
	debugFunc[102592] = medium.Leetcode_Maximize_Greatness
	debugFunc[102596] = medium.Leetcode_Check_Valid_Grid
	debugFunc[102597] = medium.Leetcode_Beautiful_Subsets
	debugFunc[102601] = medium.Leetcode_Prime_Sub_Operation
	debugFunc[102602] = medium.Leetcode_Min_Operations_I
	debugFunc[102606] = medium.Leetcode_Maximum_Cost_Substring
	debugFunc[102610] = medium.Leetcode_Find_Matrix
	// Hard
	debugFunc[100004] = hard.Leetcode_Find_Median_Sorted_Arrays
	debugFunc[100023] = hard.Leetcode_Merge_K_Lists
	debugFunc[100032] = hard.Leetcode_Longest_Valid_Parentheses
	debugFunc[100072] = hard.Leetcode_Min_Distance
	debugFunc[100087] = hard.Leetcode_Is_Scramble
	debugFunc[100502] = hard.Leetcode_Find_Maximized_Capital
	debugFunc[100768] = hard.Leetcode_Max_Chunks_To_Sorted_II
	debugFunc[101345] = hard.Leetcode_Min_Jumps
	debugFunc[101402] = hard.Leetcode_Max_Satisfaction
	debugFunc[101444] = hard.Leetcode_Ways
	debugFunc[101675] = hard.Leetcode_Minimum_Deviation
	debugFunc[101857] = hard.Leetcode_Largest_Path_Value
	debugFunc[101944] = hard.Leetcode_Can_See_Persons_Count
	debugFunc[102218] = hard.Leetcode_Max_Value_Of_Coins
	debugFunc[102306] = hard.Leetcode_Distinct_Names
	debugFunc[102360] = hard.Leetcode_Longest_Cycle
	debugFunc[102444] = hard.Leetcode_Count_Subarrays
	debugFunc[102547] = hard.Leetcode_Min_Cost
	// Other
	debugFunc[100585] = Leetcode_SQL // Medium
	debugFunc[100601] = Leetcode_SQL // Hard
	//Contest
}

func Leetcode_Debug(problem int) {
	if Invoke, hasSolution := debugFunc[normalize+problem]; hasSolution {
		Invoke()
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

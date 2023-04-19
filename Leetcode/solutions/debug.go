package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/solutions/easy"
	"leetcode.com/Leetcode/solutions/hard"
	"leetcode.com/Leetcode/solutions/medium"
)

type Solution map[int]func()

// Format problem th to int
const Normalize int = 100000

// problem th convert to problem + normalize
var _SOLUTIONS_ Solution

func init() {
	_SOLUTIONS_ = make(map[int]func())
	// Easy
	_SOLUTIONS_[100001] = easy.Leetcode_Two_Sum
	_SOLUTIONS_[100009] = easy.LeetCode_Is_Palindrome
	_SOLUTIONS_[100014] = easy.Leetcode_Longest_Common_Prefix
	_SOLUTIONS_[100020] = easy.Leetcode_Is_Valid
	_SOLUTIONS_[100021] = easy.Leetcode_Merge_Two_Lists
	_SOLUTIONS_[100026] = easy.Leetcode_Remove_Duplicates
	_SOLUTIONS_[100027] = easy.Leetcode_Remove_Element
	_SOLUTIONS_[100028] = easy.Leetcode_StrStr
	_SOLUTIONS_[100035] = easy.Leetcode_Search_Insert
	_SOLUTIONS_[100058] = easy.Leetcode_Length_Of_Last_Word
	_SOLUTIONS_[100067] = easy.Leetcode_Add_Binary
	_SOLUTIONS_[100101] = easy.Leetcode_Is_Symmetric
	_SOLUTIONS_[100104] = easy.LeetCode_Max_Depth
	_SOLUTIONS_[100121] = easy.LeetCode_Max_Profit
	_SOLUTIONS_[100219] = easy.Leetcode_Contains_Near_By_Duplicate
	_SOLUTIONS_[100226] = easy.Leetcode_Invert_Tree
	_SOLUTIONS_[100290] = easy.Leetcode_Word_Pattern
	_SOLUTIONS_[100509] = easy.Leetcode_Fibonacci
	_SOLUTIONS_[100605] = easy.Leetcode_Can_Place_Flowers
	_SOLUTIONS_[100704] = easy.Leetcode_Search
	_SOLUTIONS_[100783] = easy.Leetcode_Min_Diff_In_BST
	_SOLUTIONS_[100925] = easy.Leetcode_Is_Long_PressedName
	_SOLUTIONS_[100953] = easy.Leetcode_Is_Alien_Sorted
	_SOLUTIONS_[100989] = easy.Leetcode_Add_To_Array_Form
	_SOLUTIONS_[101047] = easy.Leetcode_Remove_Adjacent_Duplicates
	_SOLUTIONS_[101071] = easy.Leetcode_Gcd_Of_Strings
	_SOLUTIONS_[101137] = easy.Leetcode_Tribonacci
	_SOLUTIONS_[101207] = easy.Leetcode_Unique_Occurrences
	_SOLUTIONS_[101431] = easy.Leetcode_Kids_With_Candies
	_SOLUTIONS_[101470] = easy.Leetcode_Shuffle
	_SOLUTIONS_[101539] = easy.Leetcode_Find_Kth_Positive
	_SOLUTIONS_[101572] = easy.Leetcode_Diagonal_Sum
	_SOLUTIONS_[101768] = easy.Leetcode_Merge_Alternately
	_SOLUTIONS_[101925] = easy.Leetcode_Count_Triples
	_SOLUTIONS_[102399] = easy.Leetcode_Check_Distances
	_SOLUTIONS_[102566] = easy.Leetcode_Min_Max_Difference
	_SOLUTIONS_[102570] = easy.Leetcode_Merge_Arrays
	_SOLUTIONS_[102574] = easy.Leetcode_Left_Rigth_Difference
	_SOLUTIONS_[102578] = easy.Leetcode_Split_Num
	_SOLUTIONS_[102582] = easy.Leetcode_Pass_The_Pillow
	_SOLUTIONS_[102586] = easy.Leetcode_Vowel_Strings
	_SOLUTIONS_[102591] = easy.Leetcode_Dist_Money
	_SOLUTIONS_[102595] = easy.Leetcode_Even_Odd_Bit
	_SOLUTIONS_[102600] = easy.Leetcode_K_Items_With_Maximum_Sum
	_SOLUTIONS_[102605] = easy.Leetcode_Min_Number
	_SOLUTIONS_[102609] = easy.Leetcode_Find_The_Longest_Balanced_Substring
	_SOLUTIONS_[102639] = easy.Leetcode_Max_Value_Of_Coins
	_SOLUTIONS_[102643] = easy.Leetcode_Row_And_Maximum_Ones
	_SOLUTIONS_[102644] = easy.Leetcode_Max_Div_Score
	// Medium
	_SOLUTIONS_[100002] = medium.Leetcode_Add_Two_Numbers
	_SOLUTIONS_[100003] = medium.Leetcode_Length_Of_Longest_Substring
	_SOLUTIONS_[100005] = medium.Leetcode_Longest_Palindrome
	_SOLUTIONS_[100015] = medium.Leetcode_Three_Sum
	_SOLUTIONS_[100045] = medium.Leetcode_Jumb_Game_II
	_SOLUTIONS_[100064] = medium.Leetcode_Min_Path_Sum
	_SOLUTIONS_[100071] = medium.Leetcode_Simplify_Path
	_SOLUTIONS_[100103] = medium.Leetcode_Zigzag_Level_Order
	_SOLUTIONS_[100106] = medium.Leetcode_Build_Trees
	_SOLUTIONS_[100109] = medium.Leetcode_Sorted_List_To_BST
	_SOLUTIONS_[100129] = medium.Leetcode_Sum_Numbers
	_SOLUTIONS_[100133] = medium.Leetcode_Clone_Graph
	_SOLUTIONS_[100142] = medium.Leetcode_Detect_Cycle
	_SOLUTIONS_[100198] = medium.Leetcode_House_Robber
	_SOLUTIONS_[100208] = medium.Leetcode_Trie_Constructor
	_SOLUTIONS_[100211] = medium.Leetcode_Word_Dictionary
	_SOLUTIONS_[100299] = medium.Leetcode_Get_Hint
	_SOLUTIONS_[100382] = medium.Leetcode_Constructor
	_SOLUTIONS_[100427] = medium.Leetcode_Construct
	_SOLUTIONS_[100443] = medium.Leetcode_Compress
	_SOLUTIONS_[100516] = medium.Leetcode_Longest_Palindrome_Subseq
	_SOLUTIONS_[100540] = medium.Leetcode_Single_Non_Duplicate
	_SOLUTIONS_[100567] = medium.Leetcode_Check_Inclusion
	_SOLUTIONS_[100652] = medium.Leetcode_Find_Duplicate_Subtrees
	_SOLUTIONS_[100769] = medium.Leetcode_Max_Chunks_To_Sorted
	_SOLUTIONS_[100881] = medium.Leetcode_Num_Rescue_Boats
	_SOLUTIONS_[100875] = medium.Leetcode_Min_Eating_Speed
	_SOLUTIONS_[100912] = medium.Leetcode_Sort_Array
	_SOLUTIONS_[100946] = medium.Leetcode_Validate_Stack_Sequences
	_SOLUTIONS_[100958] = medium.Leetcode_Is_Complete_Tree
	_SOLUTIONS_[100983] = medium.Leetcode_Min_Cost_Tickets
	_SOLUTIONS_[101011] = medium.Leetcode_Ship_Within_Days
	_SOLUTIONS_[101020] = medium.Leetcode_Num_Enclaves
	_SOLUTIONS_[101049] = medium.Leetcode_Last_Stone_Weight_II
	_SOLUTIONS_[101129] = medium.Leetcode_Shortest_Alternating_Paths
	_SOLUTIONS_[101162] = medium.Leetcode_Max_Distance
	_SOLUTIONS_[101254] = medium.Leetcode_Closed_Island
	_SOLUTIONS_[101319] = medium.Leetcode_Make_Connected
	_SOLUTIONS_[101348] = medium.Leetcode_Tweet_Counts_Per_Frequency
	_SOLUTIONS_[101372] = medium.Leetcode_Longest_ZigZag
	_SOLUTIONS_[101466] = medium.Leetcode_Min_Reorder
	_SOLUTIONS_[101472] = medium.Leetcode_Design_Browser_History
	_SOLUTIONS_[101584] = medium.Leetcode_minCost_Connect_Points
	_SOLUTIONS_[101954] = medium.Leetcode_Minimum_Perimeter
	_SOLUTIONS_[102145] = medium.Leetcode_Number_Of_Arrays
	_SOLUTIONS_[102187] = medium.Leetcode_Minimum_Time
	_SOLUTIONS_[102300] = medium.Leetcode_Successful_Pairs
	_SOLUTIONS_[102316] = medium.Leetcode_Count_Pairs
	_SOLUTIONS_[102348] = medium.Leetcode_Zero_Filled_Subarray
	_SOLUTIONS_[102390] = medium.Leetcode_Remove_Stars
	_SOLUTIONS_[102405] = medium.Leetcode_Partition_String
	_SOLUTIONS_[102439] = medium.Leetcode_Minimize_Array_Value
	_SOLUTIONS_[102477] = medium.Leetcode_Minimum_Fuel_Cost
	_SOLUTIONS_[102492] = medium.Leetcode_Min_Score
	_SOLUTIONS_[102567] = medium.Leetcode_Minimize_Sum
	_SOLUTIONS_[102571] = medium.Leetcode_Min_Operations
	_SOLUTIONS_[102575] = medium.Leetcode_Divisibility_Array
	_SOLUTIONS_[102576] = medium.Leetcode_Max_Num_Of_Marked_Indices
	_SOLUTIONS_[102579] = medium.Leetcode_Colored_Cells
	_SOLUTIONS_[102580] = medium.Leetcode_Count_Ways
	_SOLUTIONS_[102583] = medium.Leetcode_Kth_Largest_Level_Sum
	_SOLUTIONS_[102587] = medium.Leetcode_Max_Score
	_SOLUTIONS_[102588] = medium.Leetcode_Beautiful_Subarrays
	_SOLUTIONS_[102592] = medium.Leetcode_Maximize_Greatness
	_SOLUTIONS_[102596] = medium.Leetcode_Check_Valid_Grid
	_SOLUTIONS_[102597] = medium.Leetcode_Beautiful_Subsets
	_SOLUTIONS_[102601] = medium.Leetcode_Prime_Sub_Operation
	_SOLUTIONS_[102602] = medium.Leetcode_Min_Operations_I
	_SOLUTIONS_[102606] = medium.Leetcode_Maximum_Cost_Substring
	_SOLUTIONS_[102610] = medium.Leetcode_Find_Matrix
	_SOLUTIONS_[102640] = medium.Leetcode_Find_Prefix_Score
	_SOLUTIONS_[102641] = medium.Leetcode_Replace_Value_In_Tree
	_SOLUTIONS_[102645] = medium.Leetcode_Add_Minimum
	// Hard
	_SOLUTIONS_[100004] = hard.Leetcode_Find_Median_Sorted_Arrays
	_SOLUTIONS_[100023] = hard.Leetcode_Merge_K_Lists
	_SOLUTIONS_[100032] = hard.Leetcode_Longest_Valid_Parentheses
	_SOLUTIONS_[100072] = hard.Leetcode_Min_Distance
	_SOLUTIONS_[100087] = hard.Leetcode_Is_Scramble
	_SOLUTIONS_[100502] = hard.Leetcode_Find_Maximized_Capital
	_SOLUTIONS_[100768] = hard.Leetcode_Max_Chunks_To_Sorted_II
	_SOLUTIONS_[101345] = hard.Leetcode_Min_Jumps
	_SOLUTIONS_[101402] = hard.Leetcode_Max_Satisfaction
	_SOLUTIONS_[101444] = hard.Leetcode_Ways
	_SOLUTIONS_[101675] = hard.Leetcode_Minimum_Deviation
	_SOLUTIONS_[101639] = hard.Leetcode_Num_Ways
	_SOLUTIONS_[101857] = hard.Leetcode_Largest_Path_Value
	_SOLUTIONS_[101944] = hard.Leetcode_Can_See_Persons_Count
	_SOLUTIONS_[102218] = hard.Leetcode_Max_Value_Of_Coins
	_SOLUTIONS_[102306] = hard.Leetcode_Distinct_Names
	_SOLUTIONS_[102360] = hard.Leetcode_Longest_Cycle
	_SOLUTIONS_[102444] = hard.Leetcode_Count_Subarrays
	_SOLUTIONS_[102547] = hard.Leetcode_Min_Cost
	_SOLUTIONS_[102646] = hard.Leetcode_Minimum_Total_Price
	// Other
	_SOLUTIONS_[100585] = Leetcode_SQL // Medium
	_SOLUTIONS_[100601] = Leetcode_SQL // Hard
	// Contest
	// _SOLUTIONS_[106378] = contest.Leetcode_Minimum_Total_Price
}

func Leetcode_Debug(problem int) {
	if Invoke, hasSolution := _SOLUTIONS_[Normalize+problem]; hasSolution {
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

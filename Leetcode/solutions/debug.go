package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/solutions/easy"
	"leetcode.com/Leetcode/solutions/hard"
	"leetcode.com/Leetcode/solutions/medium"
	"leetcode.com/Leetcode/types"
)

type Solution types.Solution

// problem th convert to problem + normalize
var _SOLUTIONS_ Solution = make(Solution)

func init() {
	// Easy
	for index, function := range easy.Solutions {
		_SOLUTIONS_[index] = function
	}

	// Medium
	for index, function := range medium.Solutions {
		_SOLUTIONS_[index] = function
	}

	// Hard
	for index, function := range hard.Solutions {
		_SOLUTIONS_[index] = function
	}

	// #region Solution
	// _SOLUTIONS_[101965] = easy.Leetcode_SQL
	// _SOLUTIONS_[102606] = medium.Leetcode_Maximum_Cost_Substring
	// _SOLUTIONS_[102619] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102620] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102621] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102626] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102629] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102634] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102635] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102637] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102648] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102665] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102666] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102667] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102677] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102695] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102703] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102704] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102715] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102723] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102724] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102725] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102726] = easy.Leetcode_Javascript
	// _SOLUTIONS_[102727] = easy.Leetcode_Javascript

	// _SOLUTIONS_[100570] = medium.Leetcode_SQL
	// _SOLUTIONS_[100585] = medium.Leetcode_SQL
	// _SOLUTIONS_[100649] = medium.Leetcode_Predict_Party_Victory
	// _SOLUTIONS_[100652] = medium.Leetcode_Find_Duplicate_Subtrees
	// _SOLUTIONS_[100662] = medium.Leetcode_Width_Of_Binary_Tree
	// _SOLUTIONS_[100714] = medium.Leetcode_Max_Profit_II
	// _SOLUTIONS_[100735] = medium.Leetcode_Asteroid_Collision
	// _SOLUTIONS_[100769] = medium.Leetcode_Max_Chunks_To_Sorted
	// _SOLUTIONS_[100785] = medium.Leetcode_Is_Bipartite
	// _SOLUTIONS_[100837] = medium.Leetcode_New_21_Game
	// _SOLUTIONS_[100849] = medium.Leetcode_Max_Dist_To_Closest
	// _SOLUTIONS_[100875] = medium.Leetcode_Min_Eating_Speed
	// _SOLUTIONS_[100881] = medium.Leetcode_Num_Rescue_Boats
	// _SOLUTIONS_[100904] = medium.Leetcode_Total_Fruit
	// _SOLUTIONS_[100912] = medium.Leetcode_Sort_Array
	// _SOLUTIONS_[100934] = medium.Leetcode_Shortest_Bridge
	// _SOLUTIONS_[100946] = medium.Leetcode_Validate_Stack_Sequences
	// _SOLUTIONS_[100958] = medium.Leetcode_Is_Complete_Tree
	// _SOLUTIONS_[100983] = medium.Leetcode_Min_Cost_Tickets
	// _SOLUTIONS_[101004] = medium.Leetcode_Longest_Ones
	// _SOLUTIONS_[101011] = medium.Leetcode_Ship_Within_Days
	// _SOLUTIONS_[101020] = medium.Leetcode_Num_Enclaves
	// _SOLUTIONS_[101027] = medium.Leetcode_Longest_Arith_Seq_Length
	// _SOLUTIONS_[101035] = medium.Leetcode_Max_Uncrossed_Lines
	// _SOLUTIONS_[101049] = medium.Leetcode_Last_Stone_Weight_II
	// _SOLUTIONS_[101091] = medium.Leetcode_Shortest_Path_Binary_Matrix
	// _SOLUTIONS_[101129] = medium.Leetcode_Shortest_Alternating_Paths
	// _SOLUTIONS_[101140] = medium.Leetcode_Stone_Game_II
	// _SOLUTIONS_[101143] = medium.Leetcode_Longest_Common_Subsequence
	// _SOLUTIONS_[101146] = medium.Leetcode_Snapshot_Array
	// _SOLUTIONS_[101161] = medium.Leetcode_Max_LevelSum
	// _SOLUTIONS_[101162] = medium.Leetcode_Max_Distance
	// _SOLUTIONS_[101201] = medium.Leetcode_Is_Ugly_III
	// _SOLUTIONS_[101254] = medium.Leetcode_Closed_Island
	// _SOLUTIONS_[101318] = medium.Leetcode_Min_Flips
	// _SOLUTIONS_[101319] = medium.Leetcode_Make_Connected
	// _SOLUTIONS_[101348] = medium.Leetcode_Tweet_Counts_Per_Frequency
	// _SOLUTIONS_[101372] = medium.Leetcode_Longest_ZigZag
	// _SOLUTIONS_[101376] = medium.Leetcode_Num_Of_Minutes
	// _SOLUTIONS_[101396] = medium.Leetcode_Design_Underground_System
	// _SOLUTIONS_[101448] = medium.Leetcode_Good_Nodes
	// _SOLUTIONS_[101456] = medium.Leetcode_Max_Vowels
	// _SOLUTIONS_[101466] = medium.Leetcode_Min_Reorder
	// _SOLUTIONS_[101472] = medium.Leetcode_Design_Browser_History
	// _SOLUTIONS_[101493] = medium.Leetcode_Longest_Subarray
	// _SOLUTIONS_[101498] = medium.Leetcode_Num_Subseq
	// _SOLUTIONS_[101514] = medium.Leetcode_Max_Probability
	// _SOLUTIONS_[101557] = medium.Leetcode_Find_Smallest_Set_Of_Vertices
	// _SOLUTIONS_[101584] = medium.Leetcode_minCost_Connect_Points
	// _SOLUTIONS_[101657] = medium.Leetcode_Close_Strings
	// _SOLUTIONS_[101679] = medium.Leetcode_Max_Operations
	// _SOLUTIONS_[101721] = medium.Leetcode_Swap_Nodes
	// _SOLUTIONS_[101802] = medium.Leetcode_Max_Value
	// _SOLUTIONS_[101954] = medium.Leetcode_Minimum_Perimeter
	// _SOLUTIONS_[102024] = medium.Leetcode_Max_Consecutive_Answers
	// _SOLUTIONS_[102038] = medium.Leetcode_Winner_Of_Game
	// _SOLUTIONS_[102044] = medium.Leetcode_Count_Max_Or_Subsets
	// _SOLUTIONS_[102090] = medium.Leetcode_Get_Averages
	// _SOLUTIONS_[102095] = medium.Leetcode_Delete_Middle
	// _SOLUTIONS_[102101] = medium.Leetcode_Maximum_Detonation
	// _SOLUTIONS_[102130] = medium.Leetcode_Pair_Sum
	// _SOLUTIONS_[102140] = medium.Leetcode_Most_Points
	// _SOLUTIONS_[102145] = medium.Leetcode_Number_Of_Arrays
	// _SOLUTIONS_[102187] = medium.Leetcode_Minimum_Time
	// _SOLUTIONS_[102275] = medium.Leetcode_Largest_Combination
	// _SOLUTIONS_[102300] = medium.Leetcode_Successful_Pairs
	// _SOLUTIONS_[102305] = medium.Leetcode_Distribute_Cookies
	// _SOLUTIONS_[102316] = medium.Leetcode_Count_Pairs
	// _SOLUTIONS_[102336] = medium.Leetcode_Smallest_Infinite_Set
	// _SOLUTIONS_[102348] = medium.Leetcode_Zero_Filled_Subarray
	// _SOLUTIONS_[102352] = medium.Leetcode_Equal_Pairs
	// _SOLUTIONS_[102390] = medium.Leetcode_Remove_Stars
	// _SOLUTIONS_[102405] = medium.Leetcode_Partition_String
	// _SOLUTIONS_[102419] = medium.Leetcode_Longest_Subarray_II
	// _SOLUTIONS_[102439] = medium.Leetcode_Minimize_Array_Value
	// _SOLUTIONS_[102462] = medium.Leetcode_Total_Cost
	// _SOLUTIONS_[102466] = medium.Leetcode_Count_Good_Strings
	// _SOLUTIONS_[102477] = medium.Leetcode_Minimum_Fuel_Cost
	// _SOLUTIONS_[102492] = medium.Leetcode_Min_Score
	// _SOLUTIONS_[102542] = medium.Leetcode_Maximum_Subsequence_Score
	// _SOLUTIONS_[102567] = medium.Leetcode_Minimize_Sum
	// _SOLUTIONS_[102571] = medium.Leetcode_Min_Operations
	// _SOLUTIONS_[102575] = medium.Leetcode_Divisibility_Array
	// _SOLUTIONS_[102576] = medium.Leetcode_Max_Num_Of_Marked_Indices
	// _SOLUTIONS_[102579] = medium.Leetcode_Colored_Cells
	// _SOLUTIONS_[102580] = medium.Leetcode_Count_Ways
	// _SOLUTIONS_[102583] = medium.Leetcode_Kth_Largest_Level_Sum
	// _SOLUTIONS_[102587] = medium.Leetcode_Max_Score
	// _SOLUTIONS_[102588] = medium.Leetcode_Beautiful_Subarrays
	// _SOLUTIONS_[102592] = medium.Leetcode_Maximize_Greatness
	// _SOLUTIONS_[102596] = medium.Leetcode_Check_Valid_Grid
	// _SOLUTIONS_[102597] = medium.Leetcode_Beautiful_Subsets
	// _SOLUTIONS_[102601] = medium.Leetcode_Prime_Sub_Operation
	// _SOLUTIONS_[102602] = medium.Leetcode_Min_Operations_I
	// _SOLUTIONS_[102610] = medium.Leetcode_Find_Matrix
	// _SOLUTIONS_[102618] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102622] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102623] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102625] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102627] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102628] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102631] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102632] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102633] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102636] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102640] = medium.Leetcode_Find_Prefix_Score
	// _SOLUTIONS_[102641] = medium.Leetcode_Replace_Value_In_Tree
	// _SOLUTIONS_[102645] = medium.Leetcode_Add_Minimum
	// _SOLUTIONS_[102649] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102653] = medium.Leetcode_Get_Subarray_Beauty
	// _SOLUTIONS_[102654] = medium.Leetcode_Min_Operations_Equal_1
	// _SOLUTIONS_[102657] = medium.Leetcode_Find_The_Prefix_Common_Array
	// _SOLUTIONS_[102658] = medium.Leetcode_Find_Max_Fish
	// _SOLUTIONS_[102661] = medium.Leetcode_First_Complete_Index
	// _SOLUTIONS_[102662] = medium.Leetcode_Minimum_Cost
	// _SOLUTIONS_[102671] = medium.Leetcode_Frequency_Tracker
	// _SOLUTIONS_[102672] = medium.Leetcode_Color_The_Array
	// _SOLUTIONS_[102673] = medium.Leetcode_Min_Increments
	// _SOLUTIONS_[102675] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102676] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102679] = medium.Leetcode_Matrix_Sum
	// _SOLUTIONS_[102680] = medium.Leetcode_Maximum_Or
	// _SOLUTIONS_[102683] = medium.Leetcode_Does_Valid_Array_Exist
	// _SOLUTIONS_[102684] = medium.Leetcode_Max_Moves
	// _SOLUTIONS_[102685] = medium.Leetcode_Count_Complete_Components
	// _SOLUTIONS_[102693] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102694] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102698] = medium.Leetcode_Punishment_Number
	// _SOLUTIONS_[102700] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102705] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102707] = medium.Leetcode_Min_Extra_Char
	// _SOLUTIONS_[102708] = medium.Leetcode_Max_Strength
	// _SOLUTIONS_[102711] = medium.Leetcode_Difference_Of_Distinct_Values
	// _SOLUTIONS_[102712] = medium.Leetcode_Minimum_Cost_II
	// _SOLUTIONS_[102718] = medium.Leetcode_Matrix_Sum_Queries
	// _SOLUTIONS_[102721] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102722] = medium.Leetcode_Javascript
	// _SOLUTIONS_[102730] = medium.LeetCode_Longest_Semi_Repetitive_Substring
	// _SOLUTIONS_[102731] = medium.LeetCode_Sum_Distance
	// _SOLUTIONS_[102734] = medium.Leetcode_Smallest_String
	// _SOLUTIONS_[102735] = medium.Leetcode_Min_Cost
	// _SOLUTIONS_[102740] = medium.Leetcode_Find_Value_Of_Partition
	// _SOLUTIONS_[102741] = medium.Leetcode_Special_Perm

	// _SOLUTIONS_[100601] = hard.Leetcode_SQL
	// #endregion
}

func Len() int {
	return len(_SOLUTIONS_)
}

func Leetcode_Debug(problem int) {
	if Invoke, hasSolution := _SOLUTIONS_[problem]; hasSolution {
		Invoke()
	} else {
		Leetcode_Empty(problem)
	}
}

func Leetcode_Check_Golang_Solution(problem int) bool {
	_, hasSolution := _SOLUTIONS_[problem]
	return hasSolution
}

func Leetcode_Empty(problem int) {
	fmt.Printf("The problem %d hasn't been solved yet!\n", problem)
}

package solutions

import (
	"fmt"

	"leetcode.com/Leetcode/types"
	"leetcode.com/Leetcode/utils"
)

func Leetcode_Two_Sum() {
	fmt.Println("Input: nums = [2,7,11,15], target = 9")
	fmt.Println("Output:", twoSum([]int{2, 7, 11, 15}, 9))
}

func Leetcode_Unique_Occurrences() {
	fmt.Println("Input: nums = [1,2,2,1,1,3]")
	fmt.Println("Output:", uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}

func LeetCode_Max_Depth() {
	var root *types.TreeNode

	root = types.LazyNodeAll(3, types.LazyNode(9), types.LazyNodeValue(20, 15, 7))
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output:", maxDepth(root))
	fmt.Println()

	root = types.LazyNodeAll(1, nil, types.LazyNode(9))
	fmt.Println("Input: root = [1,null,2]")
	fmt.Println("Output:", maxDepth(root))
}

func LeetCode_Is_Palindrome() {
	fmt.Println("Input: x = 0")
	fmt.Println("Output:", isPalindrome(0))
}

func LeetCode_Max_Profit() {
	fmt.Println("Input: x = [2, 1, 2, 1, 0, 1, 2]")
	fmt.Println("Output:", maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func Leetcode_Max_Distance() {
	fmt.Println("Input: grid = [[1,0,0],[0,0,0],[0,0,0]]")
	fmt.Println("Output:", maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

func Leetcode_Minimum_Deviation() {
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output:", minimumDeviation([]int{1, 2, 3, 4}))
	fmt.Println("Input: nums = [4,1,5,20,3]")
	fmt.Println("Output:", minimumDeviation(utils.Slice(4, 1, 5, 20, 3)))
}

func Leetcode_Remove_Duplicates() {
	fmt.Println("Input: nums = [0,0,1,1,1,2,2,3,3,4]")
	fmt.Println("Output:", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func Leetcode_Sort_Array() {
	fmt.Println("Input: nums = [5,1,1,2,0,0]")
	fmt.Println("Output:", fmt.Sprint(sortArray([]int{5, 1, 1, 2, 0, 0})))
}

func Leetcode_Length_Of_Longest_Substring() {
	fmt.Println("Input: s = 'abcabcbb'")
	fmt.Println("Output:", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println()
	fmt.Println("Input: s = 'abcdc'")
	fmt.Println("Output:", lengthOfLongestSubstring("abcdc"))
}

func Leetcode_Count_Subarrays() {
	fmt.Println("Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5")
	fmt.Println("Output:", countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))

	fmt.Println("Input: nums = [1,1,1,1], minK = 1, maxK = 1")
	fmt.Println("Output:", countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}

func Leetcode_Colored_Cells() {
	fmt.Println("Input: n = 1")
	fmt.Println("Output:", coloredCells(1))
}

func Leetcode_Length_Of_Last_Word() {
	fmt.Println("Input: s = '   fly me   to   the moon  '")
	fmt.Println("Output:", lengthOfLastWord("   fly me   to   the moon  "))
}

func Leetcode_Min_Eating_Speed() {
	fmt.Println("Input: piles = [3,6,7,11], h = 8")
	fmt.Println("Output:", minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println()
	fmt.Println("Input: ppiles = [30,11,23,4,20], h = 6")
	fmt.Println("Output:", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}

func Leetcode_Max_Score() {
	fmt.Println("Input: nums = [2,-1,0,1,-3,3,-3]")
	fmt.Println("Output:", maxScore([]int{2, -1, 0, 1, -3, 3, -3}))
}

func Leetcode_Vowel_Strings() {
	fmt.Println("Input: words = ['hey', 'aeo', 'mu', 'ooo', 'artro'], left = 1, right = 4")
	fmt.Println("Output:", vowelStrings([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4))
}

func Leetcode_Beautiful_Subarrays() {
	fmt.Println("Input: nums = [4,3,1,2,4]")
	fmt.Println("Output:", beautifulSubarrays([]int{4, 3, 1, 2, 4}))
}

func Leetcode_Sum_Numbers() {
	fmt.Println("Input: root = [1,2,3]")
	fmt.Println("Output:", sumNumbers(types.LazyNodeAll(1, types.LazyNode(2), types.LazyNode(3))))
}

func Leetcode_StrStr() {
	fmt.Println("Input: haystack = \"sadbutsad\", needle = \"sad\"")
	fmt.Println("Output:", strStr("sadbutsad", "sad"))
}

func Leetcode_Is_Complete_Tree() {
	fmt.Println("Input: root = [1,2,3,5,null,7,8]")
	fmt.Println("Output:", isCompleteTree(types.LazyNodeAll(1, types.LazyNodeAll(2, types.LazyNode(5), nil), types.LazyNodeValue(3, 7, 8))))
}

func Leetcode_Find_Median_Sorted_Arrays() {
	fmt.Println("Input: nums1 = [1], nums2 = [2,3,4,5,6]")
	fmt.Println("Output:", findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6}))
}

func Leetcode_Build_Trees() {
	fmt.Println("Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]")
	fmt.Println("Output:")
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}).Println()
}

func Leetcode_Remove_Element() {
	fmt.Println("Input: nums = [3,2,2,3], val = 3")
	fmt.Println("Output:", removeElement([]int{3, 2, 2, 3}, 3))
}

func Leetcode_Maximize_Greatness() {
	fmt.Println("Input: nums = [42,8,75,28,35,21,13,21]")
	fmt.Println("Output:", maximizeGreatness([]int{42, 8, 75, 28, 35, 21, 13, 21}))
}

func Leetcode_Even_Odd_Bit() {
	fmt.Println("Input: n = 11")
	fmt.Println("Output:", evenOddBit(11))
}

func Leetcode_Beautiful_Subsets() {
	fmt.Println("Input: nums = [10,4,5,7,2,1], k = 3")
	fmt.Println("Output:", beautifulSubsets([]int{10, 4, 5, 7, 2, 1}, 3))
}

func Leetcode_Zero_Filled_Subarray() {
	fmt.Println("Input: nums = [0,0,0,2,0,0]")
	fmt.Println("Output:", zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
}

func Leetcode_Make_Connected() {
	fmt.Println("Input: n = 4, connections = [[0,1],[0,2],[1,2]]")
	fmt.Println("Output:", makeConnected(4, utils.S2SoSliceInt("[[0,1],[0,2],[1,2]]")))
}

func Leetcode_Min_Score() {
	fmt.Println("Input: n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]")
	fmt.Println("Output:", minScore(4, utils.S2SoSliceInt("[[1,2,9],[2,3,6],[2,4,5],[1,4,7]]")))
}

func Leetcode_K_Items_With_Maximum_Sum() {
	fmt.Println("Input: numOnes = 6, numZeros = 6, numNegOnes = 6, k = 13")
	fmt.Println("Output:", kItemsWithMaximumSum(6, 6, 6, 13))
}

func Leetcode_Min_Operations() {
	fmt.Println("Input: nums = [3,1,6,8], queries = [1,5]")
	fmt.Println("Output:", minOperations([]int{3, 1, 6, 8}, []int{1, 5}))
}

func Leetcode_Prime_Sub_Operation() {
	fmt.Println("Input: nums = [4,9,6,10]")
	fmt.Println("Output:", primeSubOperation([]int{4, 9, 6, 10}))
}

func Leetcode_Longest_Palindrome() {
	fmt.Println("Input: s = 'babad'")
	fmt.Println("Output:", longestPalindrome("babad"))
}

func Leetcode_Longest_Cycle() {
	fmt.Println("Input: edges = [3,3,4,2,3]")
	fmt.Println("Output:", longestCycle([]int{3, 3, 4, 2, 3}))
	fmt.Println("Input: edges = [2,-1,3,1]")
	fmt.Println("Output:", longestCycle([]int{2, -1, 3, 1}))
}

func Leetcode_Min_Path_Sum() {
	fmt.Println("Input: grid = [[1,3,1],[1,5,1],[4,2,1]]")
	fmt.Println("Output:", minPathSum(utils.S2SoSliceInt("[[1,3,1],[1,5,1],[4,2,1]]")))
	fmt.Println("Input: grid = [[1,2,3],[4,5,6]]")
	fmt.Println("Output:", minPathSum(utils.S2SoSliceInt("[[1,2,3],[4,5,6]]")))
}

func Leetcode_Min_Cost_Tickets() {
	fmt.Println("Input: days = [1,4,6,7,8,20], costs = [2,7,15]")
	fmt.Println("Output:", mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println("Input: days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]")
	fmt.Println("Output:", mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}

func Leetcode_Is_Alien_Sorted() {
	fmt.Println("Input: words = ['hello','leetcode'], order = 'hlabcdefgijkmnopqrstuvwxyz'")
	fmt.Println("Output:", isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}

func Leetcode_Is_Valid() {
	fmt.Println("Input: s = '()[]{}'")
	fmt.Println("Output:", isValid("()[]{}"))
}

func Leetcode_Gcd_Of_Strings() {
	fmt.Println("Input: str1 = 'ABCABC', str2 = 'ABC")
	fmt.Println("Output:", gcdOfStrings("ABCABC", "ABC"))
}

func Leetcode_Contains_Near_By_Duplicate() {
	fmt.Println("Input: nums = [1,2,3,1,2,3], k = 2")
	fmt.Println("Output:", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func Leetcode_Max_Satisfaction() {
	fmt.Println("Input: satisfaction = [-1,-8,0,5,-9]")
	fmt.Println("Output:", maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	fmt.Println("Input: satisfaction = [-1,-4,-5]")
	fmt.Println("Output:", maxSatisfaction([]int{-1, -4, -5}))
}

func Leetcode_Remove_Stars() {
	fmt.Println("Input: s = 'leet**cod*e'")
	fmt.Println("Output:", removeStars("leet**cod*e"))
	fmt.Println("Input: s = 'erase*****'")
	fmt.Println("Output:", removeStars("erase*****"))
}

func Leetcode_Max_Chunks_To_Sorted_II() {
	fmt.Println("Input: arr = [4,2,2,1,1]")
	fmt.Println("Output:", maxChunksToSorted_ii([]int{4, 2, 2, 1, 1}))
	fmt.Println("Input: arr = [0,0,1,1,1]")
	fmt.Println("Output:", maxChunksToSorted_ii([]int{0, 0, 1, 1, 1}))
}

func Leetcode_Max_Chunks_To_Sorted() {
	fmt.Println("Input: arr = [4,3,2,1,0]")
	fmt.Println("Output:", maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	fmt.Println("Input: arr = [1,0,2,3,4]")
	fmt.Println("Output:", maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}

func Leetcode_Can_See_Persons_Count() {
	fmt.Println("Input: heights = [10,6,8,5,11,9]")
	fmt.Println("Output:", canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
	fmt.Println("Input: heights = [5,1,2,3,10]")
	fmt.Println("Output:", canSeePersonsCount([]int{5, 1, 2, 3, 10}))
}

func Leetcode_Is_Scramble() {
	fmt.Println("Input: s1 = 'great', s2 = 'rgeat'")
	fmt.Println("Output:", isScramble("great", "rgeat"))
	fmt.Println("Input: s1 = 'abcde', s2 = 'caebd'")
	fmt.Println("Output:", isScramble("abcde", "caebd"))
}

func Leetcode_Number_Of_Arrays() {
	fmt.Println("Input: differences = [1,-3,4], lower = 1, upper = 6")
	fmt.Println("Output:", numberOfArrays([]int{1, -3, 4}, 1, 6))
	fmt.Println("Input: differences = [3,-4,5,1,-2], lower = -4, upper = 5")
	fmt.Println("Output:", numberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5))
}

func Leetcode_Min_Cost() {
	fmt.Println("Input: nums = [1,2,1,2,1], k = 2")
	fmt.Println("Output:", minCost([]int{1, 2, 1, 2, 1}, 2))
	fmt.Println("Input: nums = [1,2,1,2,1], k = 5")
	fmt.Println("Output:", minCost([]int{1, 2, 1, 2, 1}, 5))
}

func Leetcode_Ways() {
	fmt.Println("Input: pizza = ['A..','AAA,'...'], k = 3")
	fmt.Println("Output:", ways([]string{"A..", "AAA", "..."}, 3))
	fmt.Println("Input: pizza = ['A..','AA.,'...'], k = 3")
	fmt.Println("Output:", ways([]string{"A..", "AA.", "..."}, 3))
}

func Leetcode_Search() {
	fmt.Println("Input: nums = [-1,0,3,5,9,12], target = 9")
	fmt.Println("Output:", search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func Leetcode_Zigzag_Level_Order() {
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output:", zigzagLevelOrder(types.LazyNodeAll(3, types.LazyNode(9), types.LazyNodeValue(20, 15, 7))))
}

func Leetcode_Min_Number() {
	fmt.Println("Input: nums1 = [4,1,3], nums2 = [5,7]")
	fmt.Println("Output:", minNumber([]int{1, 3}, []int{5, 7}))
}

func Leetcode_Maximum_Cost_Substring() {
	fmt.Println("Input: s = 'adaa', chars = 'd', vals = [-1000]")
	fmt.Println("Output:", maximumCostSubstring("adaa", "d", []int{-1000}))
	fmt.Println("Input: s = 'abc', chars = 'abc', vals = [-1,-1,-1]")
	fmt.Println("Output:", maximumCostSubstring("abc", "abc", []int{-1, -1, -1}))
}

func Leetcode_Find_The_Longest_Balanced_Substring() {
	fmt.Println("Input: s = '01011'")
	fmt.Println("Output:", findTheLongestBalancedSubstring("01011"))
}

func Leetcode_Find_Matrix() {
	fmt.Println("Input: nums = [1,3,4,1,2,3,1]")
	fmt.Println("Output:", findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output:", findMatrix([]int{1, 2, 3, 4}))
}

func Leetcode_Longest_Valid_Parentheses() {
	fmt.Println("Input: s = '(()'")
	fmt.Println("Output:", longestValidParentheses("(()"))
	fmt.Println("Input: s = ')()())'")
	fmt.Println("Output:", longestValidParentheses(")()())"))
}

func Leetcode_Successful_Pairs() {
	fmt.Println("Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7")
	fmt.Println("Output:", successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	fmt.Println("Input: spells = [3,1,2], potions = [8,5,8], success = 16")
	fmt.Println("Output:", successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
}

func Leetcode_Distinct_Names() {
	fmt.Println("Input: ideas = ['coffee','donuts','time','toffee']")
	fmt.Println("Output:", distinctNames([]string{"coffee", "donuts", "time", "toffee"}))
	fmt.Println("Input: ideas = ['aaa','baa', 'caa', 'bbb', 'cbb', 'dbb']")
	fmt.Println("Output:", distinctNames([]string{"aaa", "baa", "caa", "bbb", "cbb", "dbb"}))
}

func Leetcode_Jumb_Game_II() {
	fmt.Println("Input: nums = [2,3,1,1,4]")
	fmt.Println("Output:", jump_ii([]int{2, 3, 1, 1, 4}))
	fmt.Println("Input: nums = [2,3,0,1,4]]")
	fmt.Println("Output:", jump_ii([]int{2, 3, 0, 1, 4}))
}

func Leetcode_Num_Rescue_Boats() {
	fmt.Println("Input: people = [3,2,3,2,2], limit = 6")
	fmt.Println("Output:", numRescueBoats([]int{3, 2, 3, 2, 2}, 6))
	fmt.Println("Input: people = [5,1,4,2], limit = 6")
	fmt.Println("Output:", numRescueBoats([]int{5, 1, 4, 2}, 6))
}

func Leetcode_Count_Pairs() {
	fmt.Println("Input: n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]")
	fmt.Println("Output:", countPairs(7, utils.S2SoSliceInt("[[0,2],[0,5],[2,4],[1,6],[5,4]]")))
}

func Leetcode_Detect_Cycle() {
	fmt.Println("Input: head = [3,2,0,-4], pos = 1")
	node := types.NewListNode(3, 2, 0, -4)
	pos := node.Next
	node.Next.Next.Next = pos
	fmt.Println("Output:", detectCycle(node))
}

func Leetcode_Shuffle() {
	fmt.Println("Input: nums = [1,2,3,4,4,3,2,1], n = 4")
	fmt.Println("Output:", shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
}

func Leetcode_Minimum_Perimeter() {
	fmt.Println("Input: neededApples = 1000000000")
	fmt.Println("Output:", minimumPerimeter(1000000000))
}

func Leetcode_Is_Long_PressedName() {
	fmt.Println("Input: name = 'pyplrz', typed = 'ppyypllr'")
	fmt.Println("Output:", isLongPressedName("pyplrz", "ppyypllr"))
}

func Leetcode_minCost_Connect_Points() {
	fmt.Println("Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]")
	fmt.Println("Output:", minCostConnectPoints(utils.S2SoSliceInt("[[0,0],[2,2],[3,10],[5,2],[7,0]]")))
	fmt.Println("Input: points = [[3,12],[-2,5],[-4,1]]")
	fmt.Println("Output:", minCostConnectPoints(utils.S2SoSliceInt("[[3,12],[-2,5],[-4,1]]")))
}

func Leetcode_Partition_String() {
	fmt.Println("Input: s = 'gizfdfri'")
	fmt.Println("Output:", partitionString("gizfdfri"))
	fmt.Println("Input: s = 'hdklqkcssgxlvehva'")
	fmt.Println("Output:", partitionString("hdklqkcssgxlvehva"))
}

func Leetcode_Add_Binary() {
	fmt.Println("Input: a = '1010', b = '1011'")
	fmt.Println("Output:", addBinary("1010", "1011"))
}

func Leetcode_Kth_Largest_Level_Sum() {
	fmt.Println("Input: root = [5,8,9,2,1,3,7,4,6], k = 2")
	fmt.Println("Output:", kthLargestLevelSum(types.LazyNodeAll(5, types.LazyNodeAll(8, types.LazyNodeValue(2, 4, 6), types.LazyNode(1)), types.LazyNodeValue(9, 3, 7)), 2))
}

func Leetcode_Shortest_Alternating_Paths() {
	fmt.Println("Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]")
	fmt.Println("Output:", shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))
}

func Leetcode_Minimize_Array_Value() {
	fmt.Println("Input: nums = [3,7,1,6]")
	fmt.Println("Output:", minimizeArrayValue([]int{3, 7, 1, 6}))
	fmt.Println("Input: nums = [13,13,20,0,8,9,9]")
	fmt.Println("Output:", minimizeArrayValue([]int{13, 13, 20, 0, 8, 9, 9}))
}

func Leetcode_Search_Insert() {
	fmt.Println("Input: nums = [1,3,5,6], target = 7")
	fmt.Println("Output:", searchInsert([]int{1, 3, 5, 6}, 7))
}

func Leetcode_Merge_Two_Lists() {
	fmt.Println("Input: list1 = [1,2,4], list2 = [1,3,4]")
	fmt.Println("Output:", mergeTwoLists(types.NewListNode(1, 2, 4), types.NewListNode(1, 3, 4)))
}

func Leetcode_Word_Pattern() {
	fmt.Println("Input: pattern = 'abba', s = 'dog cat cat dog'")
	fmt.Println("Output:", wordPattern("abba", "dog cat cat dog"))
}

func Leetcode_Closed_Island() {
	fmt.Println("Input: grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]")
	fmt.Println("Output:", closedIsland(utils.S2SoSliceInt("[[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]")))
	fmt.Println("Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]")
	fmt.Println("Output:", closedIsland(utils.S2SoSliceInt("[[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]")))
}

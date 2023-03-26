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
	fmt.Println("Output: ", maxDepth(root))
	fmt.Println()

	root = types.LazyNodeAll(1, nil, types.LazyNode(9))
	fmt.Println("Input: root = [1,null,2]")
	fmt.Println("Output: ", maxDepth(root))
}

func LeetCode_Is_Palindrome() {
	fmt.Println("Input: x = 0")
	fmt.Println("Output: ", isPalindrome(0))
}

func LeetCode_Max_Profit() {
	fmt.Println("Input: x = [2, 1, 2, 1, 0, 1, 2]")
	fmt.Println("Output: ", maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func Leetcode_Max_Distance() {
	fmt.Println("Input: grid = \"[[1,0,0],[0,0,0],[0,0,0]]\"")
	fmt.Println("Output: ", maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

func Leetcode_Minimum_Deviation() {
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output: ", minimumDeviation([]int{1, 2, 3, 4}))
	fmt.Println()
	fmt.Println("Input: nums = [4,1,5,20,3]")
	fmt.Println("Output: ", minimumDeviation(utils.Slice(4, 1, 5, 20, 3)))
}

func Leetcode_Remove_Duplicates() {
	fmt.Println("Input: nums = [0,0,1,1,1,2,2,3,3,4]")
	fmt.Println("Output: ", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func Leetcode_Sort_Array() {
	fmt.Println("Input: nums = [5,1,1,2,0,0]")
	fmt.Println("Output: ", fmt.Sprint(sortArray([]int{5, 1, 1, 2, 0, 0})))
}

func Leetcode_Length_Of_Longest_Substring() {
	fmt.Println("Input: s = 'abcabcbb'")
	fmt.Println("Output: ", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println()
	fmt.Println("Input: s = 'abcdc'")
	fmt.Println("Output: ", lengthOfLongestSubstring("abcdc"))
}

func Leetcode_Count_Subarrays() {
	fmt.Println("Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5")
	fmt.Println("Output: ", countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))

	fmt.Println("Input: nums = [1,1,1,1], minK = 1, maxK = 1")
	fmt.Println("Output: ", countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}

func Leetcode_Colored_Cells() {
	fmt.Println("Input: n = 1")
	fmt.Println("Output: ", coloredCells(1))
}

func Leetcode_Length_Of_Last_Word() {
	fmt.Println("Input: s = '   fly me   to   the moon  '")
	fmt.Println("Output: ", lengthOfLastWord("   fly me   to   the moon  "))
}

func Leetcode_Min_Eating_Speed() {
	fmt.Println("Input: piles = [3,6,7,11], h = 8")
	fmt.Println("Output: ", minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println()
	fmt.Println("Input: ppiles = [30,11,23,4,20], h = 6")
	fmt.Println("Output: ", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}

func Leetcode_Max_Score() {
	fmt.Println("Input: nums = [2,-1,0,1,-3,3,-3]")
	fmt.Println("Output: ", maxScore([]int{2, -1, 0, 1, -3, 3, -3}))
}

func Leetcode_Vowel_Strings() {
	fmt.Println("Input: words = ['hey', 'aeo', 'mu', 'ooo', 'artro'], left = 1, right = 4")
	fmt.Println("Output: ", vowelStrings([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4))
}

func Leetcode_Beautiful_Subarrays() {
	fmt.Println("Input: nums = [4,3,1,2,4]")
	fmt.Println("Output: ", beautifulSubarrays([]int{4, 3, 1, 2, 4}))
}

func Leetcode_Sum_Numbers() {
	fmt.Println("Input: root = [1,2,3]")
	fmt.Println("Output: ", sumNumbers(types.LazyNodeAll(1, types.LazyNode(2), types.LazyNode(3))))
}

func Leetcode_StrStr() {
	fmt.Println("Input: haystack = \"sadbutsad\", needle = \"sad\"")
	fmt.Println("Output: ", strStr("sadbutsad", "sad"))
}

func Leetcode_Is_Complete_Tree() {
	fmt.Println("Input: root = [1,2,3,5,null,7,8]")
	fmt.Println("Output: ", isCompleteTree(types.LazyNodeAll(1, types.LazyNodeAll(2, types.LazyNode(5), nil), types.LazyNodeValue(3, 7, 8))))
}

func Leetcode_Find_Median_Sorted_Arrays() {
	fmt.Println("Input: nums1 = [1], nums2 = [2,3,4,5,6]")
	fmt.Println("Output: ", findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6}))
}

func Leetcode_Build_Trees() {
	fmt.Println("Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]")
	fmt.Println("Output: ")
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}).Println()
}

func Leetcode_Remove_Element() {
	fmt.Println("Input: nums = [3,2,2,3], val = 3")
	fmt.Println("Output: ", removeElement([]int{3, 2, 2, 3}, 3))
}

func Leetcode_Maximize_Greatness() {
	fmt.Println("Input: nums = [42,8,75,28,35,21,13,21]")
	fmt.Println("Output: ", maximizeGreatness([]int{42, 8, 75, 28, 35, 21, 13, 21}))
}

func Leetcode_Even_Odd_Bit() {
	fmt.Println("Input: n = 11")
	fmt.Println("Output: ", evenOddBit(11))
}

func Leetcode_Beautiful_Subsets() {
	fmt.Println("Input: nums = [10,4,5,7,2,1], k = 3")
	fmt.Println("Output: ", beautifulSubsets([]int{10, 4, 5, 7, 2, 1}, 3))
}

func Leetcode_Zero_Filled_Subarray() {
	fmt.Println("Input: nums = [0,0,0,2,0,0]")
	fmt.Println("Output: ", zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
}

func Test_Pattern() {
	fmt.Println("Input: nums = [0,0,0,2,0,0]")
	fmt.Println("Output:", utils.Sslice([]int{0, 0, 0, 2, 0, 0}))
	fmt.Println("Input: grid = \"[[1,0,0],[0,0,0],[0,0,0]]\"")
	fmt.Println("Output:", utils.SsoSlice([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
}

func Leetcode_Make_Connected() {
	fmt.Println("Input: n = 4, connections = [[0,1],[0,2],[1,2]]")
	fmt.Println("Output: ", makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}

func Leetcode_Min_Score() {
	fmt.Println("Input: n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]")
	fmt.Println("Output: ", minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
}

func Leetcode_K_Items_With_Maximum_Sum() {
	fmt.Println("Input: numOnes = 6, numZeros = 6, numNegOnes = 6, k = 13")
	fmt.Println("Output: ", kItemsWithMaximumSum(6, 6, 6, 13))
}

func Leetcode_Min_Operations() {
	fmt.Println("Input: nums = [3,1,6,8], queries = [1,5]")
	fmt.Println("Output: ", minOperations([]int{3, 1, 6, 8}, []int{1, 5}))
}

func Leetcode_Prime_Sub_Operation() {
	fmt.Println("Input: nums = [4,9,6,10]")
	fmt.Println("Output: ", primeSubOperation([]int{4, 9, 6, 10}))
}

func Leetcode_Longest_Cycle() {
	fmt.Println("Input: edges = [3,3,4,2,3]")
	fmt.Println("Output: ", longestCycle([]int{3, 3, 4, 2, 3}))
	fmt.Println("Input: edges = [2,-1,3,1]")
	fmt.Println("Output: ", longestCycle([]int{2, -1, 3, 1}))
}

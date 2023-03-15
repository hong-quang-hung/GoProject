package solutions

import (
	"fmt"

	T "leetcode.com/Leetcode/types"
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
	var root *T.TreeNode

	root = &T.TreeNode{Val: 3, Left: T.LazyNode(9), Right: &T.TreeNode{Val: 20, Left: T.LazyNode(15), Right: T.LazyNode(7)}}
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output: ", maxDepth(root))
	fmt.Println()

	root = &T.TreeNode{Val: 1, Left: nil, Right: T.LazyNode(9)}
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
	fmt.Println("Output: ", minimumDeviation([]int{4, 1, 5, 20, 3}))
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
	fmt.Println("Output: ", sumNumbers(T.LazyNodeAll(1, T.LazyNode(2), T.LazyNode(3))))
}

func Leetcode_StrStr() {
	fmt.Println("Input: haystack = \"sadbutsad\", needle = \"sad\"")
	fmt.Println("Output: ", strStr("sadbutsad", "sad"))
}

func Leetcode_Is_Complete_Tree() {
	fmt.Println("Input: root = [1,2,3,5,null,7,8]")
	fmt.Println("Output: ", isCompleteTree(T.LazyNodeAll(1, T.LazyNodeAll(2, T.LazyNode(5), nil), T.LazyNodeValue(3, 7, 8))))
}

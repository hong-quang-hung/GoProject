package solutions

import (
	"fmt"

	T "leetcode.com/Leetcode/types"
	U "leetcode.com/Leetcode/utils"
)

func Leetcode_Two_Sum() {
	target := 9
	nums := [...]int{2, 7, 11, 15}
	fmt.Println("Input: nums = [2,7,11,15], target = 9")
	fmt.Println("Output:", twoSum(nums[:], target))
}

func LeetCode_Max_Depth() {
	var root *T.TreeNode

	root = &T.TreeNode{Val: 3, Left: T.TreeNodeBase(9), Right: &T.TreeNode{Val: 20, Left: T.TreeNodeBase(15), Right: T.TreeNodeBase(7)}}
	fmt.Println("Input: root = [3,9,20,null,null,15,7]")
	fmt.Println("Output: ", maxDepth(root))
	fmt.Println()

	root = &T.TreeNode{Val: 1, Left: nil, Right: T.TreeNodeBase(9)}
	fmt.Println("Input: root = [1,null,2]")
	fmt.Println("Output: ", maxDepth(root))
}

func LeetCode_Is_Palindrome() {
	var x int

	x = 0
	fmt.Println("Input: x = 0")
	fmt.Println("Output: ", isPalindrome(x))
	fmt.Println()

	x = 2552
	fmt.Println("Input: x = 2552")
	fmt.Println("Output: ", isPalindrome(x))
	fmt.Println()

	x = 21120
	fmt.Println("Input: x = 21120")
	fmt.Println("Output: ", isPalindrome(x))
	fmt.Println()

	x = 5
	fmt.Println("Input: x = 5")
	fmt.Println("Output: ", isPalindrome(x))
}

func LeetCode_Max_Profit() {
	prices := []int{2, 1, 2, 1, 0, 1, 2}

	fmt.Println("Input: x = [2, 1, 2, 1, 0, 1, 2]")
	fmt.Println("Output: ", maxProfit(prices))
	fmt.Println()
}

func Leetcode_Max_Distance() {
	grid := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	fmt.Println("Input: grid = \"[[1,0,1],[0,0,0],[1,0,1]]\"")
	fmt.Println("Output: ", maxDistance(grid))

	grid = [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	fmt.Println("Input: grid = \"[[1,0,0],[0,0,0],[0,0,0]]\"")
	fmt.Println("Output: ", maxDistance(grid))

	grid = [][]int{{0, 0, 1, 1, 1}, {0, 1, 1, 0, 0}, {0, 0, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}}
	fmt.Println("Input: grid = \"[[0,0,1,1,1],[0,1,1,0,0],[0,0,1,1,0],[1,0,0,0,0],[1,1,0,0,1]]\"")
	fmt.Println("Output: ", maxDistance(grid))
	U.PrintMatrix(grid)
}

func Leetcode_Minimum_Deviation() {
	grid := []int{1, 2, 3, 4}
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output: ", minimumDeviation(grid))

	grid = []int{4, 1, 5, 20, 3}
	fmt.Println("Input: nums = [4,1,5,20,3]")
	fmt.Println("Output: ", minimumDeviation(grid))
}

func Leetcode_Remove_Duplicates() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("Input: nums = [0,0,1,1,1,2,2,3,3,4]")
	fmt.Println("Output: ", removeDuplicates(nums))
}

func Leetcode_Sort_Array() {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println("Input: nums = [5,1,1,2,0,0]")
	fmt.Println("Output: ", fmt.Sprint(sortArray(nums)))
}

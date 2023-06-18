package hard

import (
	"fmt"
	"math"
	"sort"
)

// Reference: https://leetcode.com/problems/make-array-strictly-increasing/
func Leetcode_Make_Array_Increasing() {
	fmt.Println("Input: arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]")
	fmt.Println("Output:", makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}))
	fmt.Println("Input: arr1 = [1,5,3,6,7], arr2 = [4,3,1]")
	fmt.Println("Output:", makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{4, 3, 1}))
	fmt.Println("Input: arr1 = [1,5,3,6,7], arr2 = [1,6,3,3]")
	fmt.Println("Output:", makeArrayIncreasing([]int{1, 5, 3, 6, 7}, []int{1, 6, 3, 3}))
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)

	dp := make(map[[2]int]int)
	count := makeArrayIncreasingSolved(0, -1, arr1, arr2, dp)
	if count == math.MaxInt32 {
		return -1
	}
	return count
}

func makeArrayIncreasingSolved(idx, prev int, arr1, arr2 []int, dp map[[2]int]int) int {
	if idx == len(arr1) {
		return 0
	}

	key := [2]int{idx, prev}
	if value, ok := dp[key]; ok {
		return value
	}

	count := math.MaxInt32
	if arr1[idx] > prev {
		count = makeArrayIncreasingSolved(idx+1, arr1[idx], arr1, arr2, dp)
	}

	arrIdx := findNextNum(arr2, prev)
	if arrIdx < len(arr2) {
		count = min(count, makeArrayIncreasingSolved(idx+1, arr2[arrIdx], arr1, arr2, dp)+1)
	}

	dp[key] = count
	return count
}

func findNextNum(arr []int, prev int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > prev {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

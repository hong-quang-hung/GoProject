package medium

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

func abs(x int) int {
	return utils.AbsInt(x)
}

func manhattan(p1, p2 []int) int {
	return utils.ManhattanInt(p1, p2)
}

func insertAt(arr []int, index int, value int) []int {
	return utils.InsertInt(arr, index, value)
}

func gcd(a, b int) int {
	return utils.GcdInt(a, b)
}

func gcd64(a, b int64) int64 {
	return utils.GcdInt64(a, b)
}

func compareSlice(a, b []int) bool {
	return utils.CompareSliceInt(a, b)
}

func countingSort(nums []int) {
	utils.CountingSort(nums)
}

func Leetcode_SQL() {
	fmt.Printf("This is SQL solution!\n")
}

func Leetcode_Javascript() {
	fmt.Printf("This is Javascript solution!\n")
}

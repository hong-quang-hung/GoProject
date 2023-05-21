package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/gcd-sort-of-an-array/
func Leetcode_Gcd_Sort() {
	fmt.Println("Input: nums = [7,21,3]")
	fmt.Println("Output:", gcdSort([]int{7, 21, 3}))
	fmt.Println("Input: nums = [5,2,6,2]")
	fmt.Println("Output:", gcdSort([]int{5, 2, 6, 2}))
	fmt.Println("Input: nums = [10,5,9,3,15]")
	fmt.Println("Output:", gcdSort([]int{10, 5, 9, 3, 15}))
}

func gcdSort(nums []int) bool {
	n := len(nums)
	numsSort := make([]int, n)
	copy(numsSort, nums)
	sort.Ints(numsSort)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j] && gcd(nums[i], nums[j]) == 1
	})

	fmt.Println(nums)
	fmt.Println(numsSort)
	return true
}

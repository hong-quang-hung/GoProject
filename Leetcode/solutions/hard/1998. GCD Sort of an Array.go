package hard

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/gcd-sort-of-an-array/
func init() {
	Solutions[1998] = func() {
		fmt.Println("Input: nums = [7,21,3]")
		fmt.Println("Output:", gcdSort([]int{7, 21, 3}))
		fmt.Println("Input: nums = [5,2,6,2]")
		fmt.Println("Output:", gcdSort([]int{5, 2, 6, 2}))
		fmt.Println("Input: nums = [10,5,9,3,15]")
		fmt.Println("Output:", gcdSort([]int{10, 5, 9, 3, 15}))
	}
}

func gcdSort(nums []int) bool {
	union := NewUnionFind(100001)
	for _, x := range nums {
		y := x
		for p := 2; p*p <= y; p++ {
			if y%p == 0 {
				union.UnionSet(x, p)
				for y%p == 0 {
					y /= p
				}
			}
		}

		if y != 1 {
			union.UnionSet(x, y)
		}
	}

	n := len(nums)
	temp := make([]int, n)
	copy(temp, nums)
	sort.Ints(temp)
	for i := 0; i < n; i++ {
		if union.Find(nums[i]) != union.Find(temp[i]) {
			return false
		}
	}
	return true
}

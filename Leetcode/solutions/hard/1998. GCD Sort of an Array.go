package hard

import "fmt"

func init() {
	Solutions[1998] = Leetcode_Gcd_Sort
}

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
	union := NewUnionFind(n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if gcd(nums[i], nums[j]) > 1 {
				union.UnionSet(j, i)
				union.UnionSet(i, j)
			}
		}
	}

	fmt.Println(union.Parent, union.Rank, union.Count)
	for i := 0; i < n; i++ {
		union.Parent[i] = union.Find(union.Parent[i])
	}

	fmt.Println(union.Parent, union.Rank, union.Count)
	return true
}

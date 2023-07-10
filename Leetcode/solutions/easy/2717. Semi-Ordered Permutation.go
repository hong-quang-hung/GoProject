package easy

import "fmt"

func init() {
	Solutions[2717] = Leetcode_Semi_Ordered_Permutation
}

// Reference: https://leetcode.com/problems/semi-ordered-permutation/
func Leetcode_Semi_Ordered_Permutation() {
	fmt.Println("Input: nums = [2,1,4,3]")
	fmt.Println("Output:", semiOrderedPermutation([]int{2, 1, 4, 3}))
	fmt.Println("Input: nums = [2,4,1,3]")
	fmt.Println("Output:", semiOrderedPermutation([]int{2, 4, 1, 3}))
	fmt.Println("Input: nums = [1,3,4,2,5]")
	fmt.Println("Output:", semiOrderedPermutation([]int{1, 3, 4, 2, 5}))
	fmt.Println("Input: nums = [3,2,1]")
	fmt.Println("Output:", semiOrderedPermutation([]int{3, 2, 1}))
}

func semiOrderedPermutation(nums []int) int {
	n, i := len(nums), 0
	for nums[i] != 1 {
		i++
	}

	res := i
	for j := i; j > 0; j-- {
		nums[j], nums[j-1] = nums[j-1], nums[j]
	}

	for j := n - 1; j >= 0; j-- {
		if nums[j] == n {
			break
		}
		res++
	}
	return res
}

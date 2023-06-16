package hard

import "fmt"

// Reference: https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/
func Leetcode_Num_Of_Ways() {
	fmt.Println("Input: nums = [2,1,3]")
	fmt.Println("Output:", numOfWays([]int{2, 1, 3}))
	fmt.Println("Input: nums = [3,4,5,1,2]")
	fmt.Println("Output:", numOfWays([]int{3, 4, 5, 1, 2}))
	fmt.Println("Input: nums = [1, 2, 3]")
	fmt.Println("Output:", numOfWays([]int{1, 2, 3}))
}

func numOfWays(nums []int) int {
	return 0
}

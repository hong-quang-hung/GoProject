package easy

import "fmt"

// Reference: https://leetcode.com/problems/set-mismatch/
func init() {
	Solutions[645] = func() {
		fmt.Println(`Input: nums = [1,2,2,4]`)
		fmt.Println(`Output:`, findErrorNums([]int{1, 2, 2, 4}))
		fmt.Println(`Input: nums = [1,1]`)
		fmt.Println(`Output:`, findErrorNums([]int{1, 1}))
	}
}

func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res[0] = nums[i]
			res[1] = i + 1
			return res
		}
	}
	return res
}

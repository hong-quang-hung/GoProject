package easy

import "fmt"

// Reference: https://leetcode.com/problems/sort-array-by-parity-ii/
func init() {
	Solutions[922] = func() {
		fmt.Println(`Input: nums = [4,2,5,7]`)
		fmt.Println(`Output:`, sortArrayByParityII([]int{4, 2, 5, 7}))
		fmt.Println(`Input: nums = [2,3]`)
		fmt.Println(`Output:`, sortArrayByParityII([]int{2, 3}))
	}
}

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	even, odd := 0, 0
	for i := 0; i < n; i += 2 {
		for even < n && nums[even]&1 == 1 {
			even++
		}

		for odd < n && nums[odd]&1 == 0 {
			odd++
		}

		res[i], res[i+1] = nums[even], nums[odd]
		even++
		odd++
	}
	return res
}

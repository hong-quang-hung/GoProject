package hard

import "fmt"

// Reference: https://leetcode.com/problems/minimum-replacements-to-sort-the-array/
func init() {
	Solutions[2366] = func() {
		fmt.Println(`Input: nums = [3,9,3]`)
		fmt.Println(`Output:`, minimumReplacement([]int{3, 9, 3}))
		fmt.Println(`Input: nums = [1,2,3,4,5]`)
		fmt.Println(`Output:`, minimumReplacement([]int{1, 2, 3, 4, 5}))
	}
}

func minimumReplacement(nums []int) int64 {
	res, n := int64(0), len(nums)
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			continue
		}

		numsElements := int64(nums[i]+nums[i+1]-1) / int64(nums[i+1])
		res += numsElements - 1
		nums[i] = nums[i] / int(numsElements)
	}
	return res
}

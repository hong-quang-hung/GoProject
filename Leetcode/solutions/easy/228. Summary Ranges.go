package easy

import "fmt"

// Reference: https://leetcode.com/problems/summary-ranges/
func init() {
	Solutions[228] = func() {
		fmt.Println("Input: nums = [0,1,2,4,5,7]")
		fmt.Println("Output:", summaryRanges([]int{0, 1, 2, 4, 5, 7}))
		fmt.Println("Input: nums = [0,2,3,4,6,8,9]")
		fmt.Println("Output:", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	}
}

func summaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); {
		j := i + 1
		for (j < len(nums)) && (nums[j]-nums[i] == j-i) {
			j++
		}

		if j-i == 1 {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
		}

		i = j
	}
	return res
}

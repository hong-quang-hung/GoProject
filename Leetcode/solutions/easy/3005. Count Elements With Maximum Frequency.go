package easy

import "fmt"

// Reference: https://leetcode.com/problems/count-elements-with-maximum-frequency/
func init() {
	Solutions[2709] = func() {
		fmt.Println(`Input: nums = [1,2,2,3,1,4]`)
		fmt.Println(`Output:`, maxFrequencyElements([]int{1, 2, 2, 3, 1, 4}))
		fmt.Println(`Input: nums = [1,2,3,4,5]`)
		fmt.Println(`Output:`, maxFrequencyElements([]int{1, 2, 3, 4, 5}))
	}
}

func maxFrequencyElements(nums []int) int {
	m := make([]int, len(nums)+1)
	maxVal := 0
	for _, num := range nums {
		m[num]++
		maxVal = max(maxVal, m[num])
	}

	res := 0
	for i := 1; i < len(m); i++ {
		if m[i] == maxVal {
			res += maxVal
		}
	}
	return res
}

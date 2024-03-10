package easy

import "fmt"

// Reference: https://leetcode.com/problems/intersection-of-two-arrays/
func init() {
	Solutions[349] = func() {
		fmt.Println(`Input: nums1 = [1,2,2,1], nums2 = [2,2]`)
		fmt.Println(`Output:`, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
		fmt.Println(`Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]`)
		fmt.Println(`Output:`, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make([]bool, 1001)
	for _, num := range nums1 {
		m1[num] = true
	}

	m2 := make([]bool, 1001)
	res := make([]int, 0)
	for _, num := range nums2 {
		if m1[num] && !m2[num] {
			res = append(res, num)
			m2[num] = true
		}
	}
	return res
}

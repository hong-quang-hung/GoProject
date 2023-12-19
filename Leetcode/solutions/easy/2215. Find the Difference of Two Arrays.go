package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-the-difference-of-two-arrays/
func init() {
	Solutions[2215] = func() {
		fmt.Println(`Input: nums1 = [1,2,3], nums2 = [2,4,6]`)
		fmt.Println(`Output:`, findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
		fmt.Println(`Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]`)
		fmt.Println(`Output:`, findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
	}
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for _, num := range nums1 {
		m1[num] = true
	}

	res := make([][]int, 2)
	for _, num := range nums2 {
		if !m1[num] {
			if !m2[num] {
				res[1] = append(res[1], num)
				m2[num] = true
			}
		} else {
			m2[num] = true
			delete(m1, num)
		}
	}

	for num := range m1 {
		res[0] = append(res[0], num)
	}
	return res
}

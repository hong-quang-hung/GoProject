package easy

import "fmt"

// Reference: https://leetcode.com/problems/unique-number-of-occurrences/
func init() {
	Solutions[1207] = func() {
		fmt.Println(`Input: nums = [1,2,2,1,1,3]`)
		fmt.Println(`Output:`, uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	}
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}

	s := make(map[int]bool)
	for _, v := range m {
		if _, ok := s[v]; ok {
			return false
		} else {
			s[v] = true
		}
	}
	return true
}

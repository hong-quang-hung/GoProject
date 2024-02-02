package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/sequential-digits/
func init() {
	Solutions[1291] = func() {
		fmt.Println(`Input: low = 100, high = 300`)
		fmt.Println(`Output:`, sequentialDigits(100, 300))
		fmt.Println(`Input: low = 1000, high = 13000`)
		fmt.Println(`Output:`, sequentialDigits(1000, 13000))
		fmt.Println(`Input: low = 89, high = 234`)
		fmt.Println(`Output:`, sequentialDigits(89, 234))
	}
}

func sequentialDigits(low int, high int) []int {
	res := []int{12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789, 1234, 2345, 3456, 4567, 5678, 6789, 12345, 23456, 34567, 45678, 56789, 123456, 234567, 345678, 456789, 1234567, 2345678, 3456789, 12345678, 23456789, 123456789}
	left := sort.Search(len(res), func(i int) bool { return res[i] >= low })
	right := sort.Search(len(res), func(i int) bool { return res[i] > high })
	return res[left:right]
}

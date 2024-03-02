package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/number-of-lines-to-write-string/
func init() {
	Solutions[806] = func() {
		fmt.Println(`Input: widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10], s = "abcdefghijklmnopqrstuvwxyz"`)
		fmt.Println(`Output:`, numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
		fmt.Println(`Input: widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10], s = "bbbcccdddaaa"`)
		fmt.Println(`Output:`, numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
	}
}

func numberOfLines(widths []int, s string) []int {
	num, currSum := 0, 0
	for _, ch := range s {
		currVal := widths[int((ch - 'a'))]
		if currSum+currVal > 100 {
			currSum = currVal
			num++
		} else {
			currSum += currVal
		}
	}

	if currSum > 0 {
		num++
	}
	return []int{num, currSum}
}

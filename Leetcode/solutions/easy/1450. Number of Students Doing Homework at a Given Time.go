package easy

import "fmt"

// Reference: https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time/
func init() {
	Solutions[1450] = func() {
		fmt.Println(`Input: startTime = [1,2,3], endTime = [3,2,7], queryTime = 4`)
		fmt.Println(`Output:`, busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
		fmt.Println(`Input: startTime = [4], endTime = [4], queryTime = 4`)
		fmt.Println(`Output:`, busyStudent([]int{4}, []int{4}, 4))
	}
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	return 0
}

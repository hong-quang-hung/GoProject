package hard

import "fmt"

// Reference: https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/
func init() {
	Solutions[1964] = func() {
		fmt.Println(`Input: obstacles = [1,2,3,2]`)
		fmt.Println(`Output:`, longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2}))
		fmt.Println(`Input: obstacles = [2,2,1]`)
		fmt.Println(`Output:`, longestObstacleCourseAtEachPosition([]int{2, 2, 1}))
		fmt.Println(`Input: obstacles = [3,1,5,6,4,2]`)
		fmt.Println(`Output:`, longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))
	}
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	sub, res := []int{}, []int{}
	for _, o := range obstacles {
		i := longestObstacleCourseAtEachPositionSearch(sub, o)
		if i == len(sub) {
			sub = append(sub, o)
		} else {
			sub[i] = o
		}
		res = append(res, i+1)
	}
	return res
}

func longestObstacleCourseAtEachPositionSearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

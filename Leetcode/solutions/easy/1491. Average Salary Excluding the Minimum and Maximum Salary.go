package easy

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
func init() {
	Solutions[1491] = func() {
		fmt.Println(`Input: salary = [4000,3000,1000,2000]`)
		fmt.Println(`Output:`, average([]int{4000, 3000, 1000, 2000}))
		fmt.Println(`Input: salary = [8000,9000,2000,3000,6000,1000]`)
		fmt.Println(`Output:`, average([]int{8000, 9000, 2000, 3000, 6000, 1000}))
	}
}

func average(salary []int) float64 {
	minS, maxS, total := math.MaxInt, 0, float64(0)
	for _, s := range salary {
		minS = min(minS, s)
		maxS = max(maxS, s)
		total += float64(s)
	}
	return (total - float64(maxS) - float64(minS)) / float64(len(salary)-2)
}

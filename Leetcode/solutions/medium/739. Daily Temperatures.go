package medium

import "fmt"

// Reference: https://leetcode.com/problems/daily-temperatures/
func init() {
	Solutions[739] = func() {
		fmt.Println("Input: temperatures = [73,74,75,71,69,72,76,73]")
		fmt.Println("Output:", dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
		fmt.Println("Input: temperatures = [30,40,50,60]")
		fmt.Println("Output:", dailyTemperatures([]int{30, 40, 50, 60}))
		fmt.Println("Input: temperatures = [30,60,90]")
		fmt.Println("Output:", dailyTemperatures([]int{30, 60, 90}))
	}
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	for i := 0; i < n; i++ {
	}
	return res
}

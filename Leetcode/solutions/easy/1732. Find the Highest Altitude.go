package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-the-highest-altitude/
func Leetcode_Largest_Altitude() {
	fmt.Println("Input: gain = [-5,1,5,0,-7]")
	fmt.Println("Output:", largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println("Input: gain = [-4,-3,-2,-1,4,3,2]")
	fmt.Println("Output:", largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}

func largestAltitude(gain []int) int {
	res := max(0, gain[0])
	for i := 1; i < len(gain); i++ {
		gain[i] += gain[i-1]
		res = max(res, gain[i])
	}
	return res
}

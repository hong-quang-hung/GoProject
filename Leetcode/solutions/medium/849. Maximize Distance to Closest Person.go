package medium

import "fmt"

func init() {
	Solutions[849] = Leetcode_Max_Dist_To_Closest
}

// Reference: https://leetcode.com/problems/maximize-distance-to-closest-person/
func Leetcode_Max_Dist_To_Closest() {
	fmt.Println("Input: seats = [1,0,0,0,1,0,1]")
	fmt.Println("Output:", maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
	fmt.Println("Input: seats = [1,0,0,0]")
	fmt.Println("Output:", maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	best, zeroes := 0, 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			zeroes++
			if zeroes == i+1 {
				best = max(best, zeroes)
			} else {
				best = max(best, (zeroes+1)/2)
			}
		} else {
			zeroes = 0
		}
	}
	return max(best, zeroes)
}

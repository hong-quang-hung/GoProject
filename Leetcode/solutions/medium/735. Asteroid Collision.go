package medium

import "fmt"

func init() {
	Solutions[735] = Leetcode_Asteroid_Collision
}

// Reference: https://leetcode.com/problems/asteroid-collision/
func Leetcode_Asteroid_Collision() {
	fmt.Println("Input: nums = [5,10,-5]")
	fmt.Println("Output:", asteroidCollision([]int{5, 10, -5}))
	fmt.Println("Input: nums = [8,-8]")
	fmt.Println("Output:", asteroidCollision([]int{8, -8}))
	fmt.Println("Input: nums = [-2,-1,1,2]")
	fmt.Println("Output:", asteroidCollision([]int{-2, -1, 1, 2}))
}

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	for _, asteroid := range asteroids {
		if asteroid > 0 {
			res = append([]int{asteroid}, res...)
		} else {
			for len(res) != 0 && res[0] > 0 && res[0] < -asteroid {
				res = res[1:]
			}

			if len(res) == 0 || res[0] < 0 {
				res = append([]int{asteroid}, res...)
			} else if res[0] == -asteroid {
				res = res[1:]
			}
		}
	}

	// Reverser
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

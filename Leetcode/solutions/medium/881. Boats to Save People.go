package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/boats-to-save-people/
func init() {
	Solutions[881] = func() {
		fmt.Println("Input: people = [3,2,3,2,2], limit = 6")
		fmt.Println("Output:", numRescueBoats([]int{3, 2, 3, 2, 2}, 6))
		fmt.Println("Input: people = [5,1,4,2], limit = 6")
		fmt.Println("Output:", numRescueBoats([]int{5, 1, 4, 2}, 6))
	}
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j, res := 0, len(people)-1, 0
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		res++
	}
	return res
}

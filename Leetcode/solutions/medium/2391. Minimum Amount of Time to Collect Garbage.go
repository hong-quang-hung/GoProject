package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
func init() {
	Solutions[2391] = func() {
		fmt.Println("Input: garbage = [\"G\",\"P\",\"GP\",\"GG\"], travel = [2,4,3]")
		fmt.Println("Output:", garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 3, 4}))
		fmt.Println("Input: garbage = [\"MMM\",\"PGM\",\"GP\"], travel = [3,10]")
		fmt.Println("Output:", garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))
	}
}

func garbageCollection(garbage []string, travel []int) int {
	for i := 1; i < len(travel); i++ {
		travel[i] += travel[i-1]
	}

	m := make(map[rune]int)
	res := 0
	for i, g := range garbage {
		for _, ch := range g {
			m[ch] = i
		}
		res += len(g)
	}

	for _, ch := range "GPM" {
		if last, ok := m[ch]; ok {
			if last > 0 {
				res += travel[last-1]
			}
		}
	}
	return res
}

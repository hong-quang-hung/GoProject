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
	G, P, M := 0, 0, 0
	return G + P + M
}

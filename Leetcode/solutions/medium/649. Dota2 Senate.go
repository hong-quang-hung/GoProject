package medium

import "fmt"

// Reference: https://leetcode.com/problems/dota2-senate/
func init() {
	Solutions[649] = func() {
		fmt.Println("Input: senate = 'RD'")
		fmt.Println("Output:", predictPartyVictory("RD"))
		fmt.Println("Input: senate = 'RDD'")
		fmt.Println("Output:", predictPartyVictory("RDD"))
		fmt.Println("Input: senate = 'DDRRR'")
		fmt.Println("Output:", predictPartyVictory("DDRRR"))
	}
}

func predictPartyVictory(senate string) string {
	R, D := []int{}, []int{}
	n := len(senate)
	for i, ch := range senate {
		if ch == 'R' {
			R = append(R, i)
		} else {
			D = append(D, i)
		}
	}

	for len(R) > 0 && len(D) > 0 {
		if R[0] < D[0] {
			R = append(R, R[0]+n)
		} else {
			D = append(D, D[0]+n)
		}
		R = R[1:]
		D = D[1:]
	}

	if len(D) == 0 {
		return "Radiant"
	}
	return "Dire"
}

package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/dota2-senate/
func Leetcode_Predict_Party_Victory() {
	fmt.Println("Input: senate = 'RD'")
	fmt.Println("Output:", predictPartyVictory("RD"))
	fmt.Println("Input: senate = 'RDD'")
	fmt.Println("Output:", predictPartyVictory("RDD"))
	fmt.Println("Input: senate = 'DDRRR'")
	fmt.Println("Output:", predictPartyVictory("DDRRR"))
}

func predictPartyVictory(senate string) string {
	res := []string{"Radiant", "Dire"}
	left, right := 0, len(senate)-1
	for left < right {
		left++
	}
	return res[0]
}

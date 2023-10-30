package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/poor-pigs/
func init() {
	Solutions[458] = func() {
		fmt.Println("Input: buckets = 4, minutesToDie = 15, minutesToTest = 15")
		fmt.Println("Output:", poorPigs(4, 15, 15))
		fmt.Println("Input: buckets = 4, minutesToDie = 15, minutesToTest = 30")
		fmt.Println("Output:", poorPigs(4, 15, 30))
	}
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	testRound := minutesToTest / minutesToDie
	testGroup := testRound + 1
	res := int(math.Ceil(math.Log2(float64(buckets)) / math.Log2(float64(testGroup))))
	return res
}

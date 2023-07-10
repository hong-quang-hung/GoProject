package easy

import "fmt"

func init() {
	Solutions[1523] = Leetcode_Count_Odds
}

// Reference: https://leetcode.com/problems/path-sum/
func Leetcode_Count_Odds() {
	fmt.Println("Input: low = 3, high = 7")
	fmt.Println("Output:", countOdds(3, 7))
	fmt.Println("Input: low = 8, high = 10")
	fmt.Println("Output:", countOdds(8, 10))
}

func countOdds(low int, high int) int {
	return ((high + 1) >> 1) - (low >> 1)
}

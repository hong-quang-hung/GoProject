package easy

import "fmt"

// Reference: https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
func init() {
	Solutions[1903] = func() {
		fmt.Println("Input: num = \"52\"")
		fmt.Println("Output:", largestOddNumber("52"))
		fmt.Println("Input: num = \"4206\"")
		fmt.Println("Output:", largestOddNumber("4206"))
		fmt.Println("Input: num = \"35427\"")
		fmt.Println("Output:", largestOddNumber("35427"))
	}
}

func largestOddNumber(num string) string {
	i := len(num) - 1
	for i >= 0 && (num[i]-'0')&1 == 0 {
		i--
	}
	return num[0 : i+1]
}

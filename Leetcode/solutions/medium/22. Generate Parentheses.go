package medium

import "fmt"

func init() {
	Solutions[22] = Leetcode_Generate_Parenthesis
}

// Reference: https://leetcode.com/problems/generate-parentheses/
func Leetcode_Generate_Parenthesis() {
	fmt.Println("Input: n = 3")
	fmt.Println("Output:", generateParenthesis(3))
	fmt.Println("Input: n = 1")
	fmt.Println("Output:", generateParenthesis(1))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	res := []string{}
	for leftCount := 0; leftCount < n; leftCount++ {
		for _, leftString := range generateParenthesis(leftCount) {
			for _, rightString := range generateParenthesis(n - 1 - leftCount) {
				res = append(res, "("+leftString+")"+rightString)
			}
		}
	}
	return res
}

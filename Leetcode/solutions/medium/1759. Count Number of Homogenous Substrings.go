package medium

import "fmt"

// Reference: https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
func init() {
	Solutions[1759] = func() {
		fmt.Println("Input: s = \"abbcccaa\"")
		fmt.Println("Output:", countHomogenous("abbcccaa"))
		fmt.Println("Input: s = \"xy\"")
		fmt.Println("Output:", countHomogenous("xy"))
		fmt.Println("Input: s = \"zzzzz\"")
		fmt.Println("Output:", countHomogenous("zzzzz"))
	}
}

func countHomogenous(s string) int {
	res := 0
	fmt.Println(5 + 4 + 3 + 2 + 1)
	return res
}

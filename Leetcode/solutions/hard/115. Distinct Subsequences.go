package hard

import "fmt"

// Reference: https://leetcode.com/problems/distinct-subsequences/
func init() {
	Solutions[115] = func() {
		fmt.Println(`Input: s = "rabbbit", t = "rabbit"`)
		fmt.Println(`Output:`, numDistinct("rabbbit", "rabbit"))
		fmt.Println(`Input: s = "babgbag", t = "bag"`)
		fmt.Println(`Output:`, numDistinct("babgbag", "bag"))
	}
}

func numDistinct(s string, t string) int {
	return 0
}

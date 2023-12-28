package hard

import "fmt"

// Reference: https://leetcode.com/problems/string-compression-ii/
func init() {
	Solutions[1531] = func() {
		fmt.Println(`Input: s = "aaabcccd", k = 2`)
		fmt.Println(`Output:`, getLengthOfOptimalCompression("aaabcccd", 2))
		fmt.Println(`Input: s = "aaabcccd", k = 2`)
		fmt.Println(`Output:`, getLengthOfOptimalCompression("aaabcccd", 2))
		fmt.Println(`Input: s = "aaaaaaaaaaa", k = 0`)
		fmt.Println(`Output:`, getLengthOfOptimalCompression("aaaaaaaaaaa", 0))
	}
}

func getLengthOfOptimalCompression(s string, k int) int {
	return 0
}
